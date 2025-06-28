import 'dart:convert';

import 'package:http/http.dart' as http;
import 'package:tracker_app/constants/storage-keys.constants.dart';
import 'package:tracker_app/enums/method-type.dart';
import 'package:tracker_app/models/refresh-token-response.model.dart';
import 'package:tracker_app/services/local-storage.service.dart';

import '../../models/common/http-response.model.ts.dart';

class HttpClientBuilder {
  String? _url;
  Map<String, String> _queryParameters = {};
  bool _useAuth = true;
  Map<String, dynamic>? _body;
  bool _hasRetried = false;
  MethodType? _methodType;

  static final _storageService = LocalStorageService();

  final Map<String,String> _headers = {
    'Content-type' : 'application/json',
    'Accept': 'application/json',
  };

  final _connectionStringBase = "http://localhost:5100";

  HttpClientBuilder withUrl(String url) {
    _url = url;
    return this;
  }

  HttpClientBuilder isAnonymous() {
    _useAuth = false;
    return this;
  }

  HttpClientBuilder withQueryParams(Map<String, String> queryParameters) {
    _queryParameters = queryParameters;
    return this;
  }

  HttpClientBuilder withQueryParameter(String key, String value) {
    _queryParameters[key] = value;
    return this;
  }

  HttpClientBuilder withBody(Map<String, dynamic> body) {
    _body = body;
    return this;
  }
  
  HttpClientBuilder withMethodType(MethodType methodType) {
    _methodType = methodType;
    return this;
  }

  Future<RequestResponse<Map<String, dynamic>>> get([String? url]) {
    return _performRequest(MethodType.GET, url);
  }

  Future<RequestResponse<Map<String, dynamic>>> post([String? url]) {
    return _performRequest(MethodType.POST, url);
  }

  Future<RequestResponse<Map<String, dynamic>>> put([String? url]) {
    return _performRequest(MethodType.PUT, url);
  }

  Future<RequestResponse<Map<String, dynamic>>> delete([String? url]) {
  return _performRequest(MethodType.DELETE, url);
}

  Future<RequestResponse<Map<String, dynamic>>> execute([String? url]) {
    if(_methodType == null){
      throw Exception("Method type is required");
    }

    return _performRequest(_methodType!, url);
  }

  Future<RequestResponse<Map<String, dynamic>>> _performRequest(MethodType method, [String? url]) async {
    _url = _buildUrl(url);
    final uri = Uri.parse(_url!);
    http.Response response;

    if (_useAuth) {
      await _populateAuthTokenHeader();
    }

    try {
      switch (method) {
        case MethodType.GET:
          response = await http.get(uri, headers: _headers);
          break;
        case MethodType.POST:
          response = await http.post(uri, headers: _headers, body: jsonEncode(_body));
          break;
        case MethodType.PUT:
          response = await http.put(uri, headers: _headers, body: _body);
          break;
        case MethodType.DELETE:
          response = await http.delete(uri, headers: _headers);
          break;
      }
    } catch (e) {
      return RequestResponse.error("Network error: $e", 500);
    }

    if (response.statusCode == 401 && !_hasRetried) {
      final refreshResponse = await _handleUnauthorizedResponse();
      if (refreshResponse.isSuccess) {
        return _performRequest(method, _url);
      }
      return RequestResponse.error("Failed to refresh token", 401);
    }

    return _getJsonFromResponse(response);
  }

  String _buildUrl([String? url]) {
    if (url != null) {
      _url = url;
    }

    if (_url == null) {
      throw Exception("Url is required");
    }

    var finalUrl = "$_connectionStringBase/$_url${_buildQueryParamString()}";

    return finalUrl;
  }

  String _buildQueryParamString() {
    var query = "?";

    _queryParameters.forEach((key, value) {
      query += "$key=$value&";
    });

    return query.substring(0, query.length - 1);
  }

  RequestResponse<Map<String, dynamic>> _getJsonFromResponse(http.Response response) {
    switch (response.statusCode) {
      case 200:
        return RequestResponse.success(json.decode(response.body));
      default:
        return RequestResponse(httpStatusCode: response.statusCode, message: response.reasonPhrase);
    }
  }

  Future<RequestResponse<Map<String, dynamic>>> _handleUnauthorizedResponse() async {
    _hasRetried = true;
    var refreshEndpoint = "$_connectionStringBase/auth/refresh";

    var refreshToken = await _storageService.fetchValue(StorageKeys.refreshTokenKey);

    if (refreshToken == null) {
      return RequestResponse.error("No refresh token found", 500);
    }

    var request = {
      'refresh_token': refreshToken
    };

    var response = await http.post(
        Uri.parse(refreshEndpoint),
        body: jsonEncode(request),
        headers: _headers);

    if (response.statusCode != 200) {
      return RequestResponse.error("Failed to refresh token", response.statusCode);
    }

    var refreshResponse = RefreshTokenResponse.fromJson(json.decode(response.body));

    try {
      await _storageService.saveValue(StorageKeys.authTokenKey, refreshResponse.accessToken);
      await _storageService.saveValue(StorageKeys.refreshTokenKey, refreshResponse.refreshToken);

      return RequestResponse.success(json.decode(response.body));
    } catch (storageError) {
      return RequestResponse.error("Failed to save authentication tokens: $storageError", 500);
    }
  }

  Future _populateAuthTokenHeader() async {
    var authToken = await _storageService.fetchValue(StorageKeys.authTokenKey);
    if (authToken != null) {
      _headers['Authorization'] = "Bearer $authToken";
    }
  }
}