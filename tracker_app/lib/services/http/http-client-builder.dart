import 'dart:convert';

import 'package:flutter_secure_storage/flutter_secure_storage.dart';
import 'package:http/http.dart' as http;
import 'package:tracker_app/constants/storage-keys.constants.dart';
import 'package:tracker_app/models/refresh-token-response.model.dart';
import 'package:tracker_app/services/local-storage.service.dart';

import '../../models/common/http-response.model.ts.dart';

class HttpClientBuilder {
  String? _url;
  Map<String, String> _queryParameters = {};
  bool _useAuth = true;
  Map<String, dynamic>? _body;
  bool _hasRetried = false;

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

  Future<RequestResponse<Map<String, dynamic>>> get([String? url]) async {
    _url = _buildUrl(url);

    final response = await http.get(Uri.parse(_url!), headers: _headers);

    if (response.statusCode == 401 && !_hasRetried) {
      var refreshResponse = await _handleUnauthorizedResponse();
      if (refreshResponse.isSuccess) {
        return get(_url);
      }
      return RequestResponse.error("Failed to refresh token", 401);
    }

    return _getJsonFromResponse(response);
  }

  Future<RequestResponse<Map<String, dynamic>>> post([String? url]) async {
    _url = _buildUrl(url);

    var response = await http.post(
        Uri.parse(_url!),
        body: jsonEncode(_body),
        headers: _headers);

    if (response.statusCode == 401 && !_hasRetried) {
      var refreshResponse = await _handleUnauthorizedResponse();
      if (refreshResponse.isSuccess) {
        return post(_url);
      }
      return RequestResponse.error("Failed to refresh token", 401);
    }

    return _getJsonFromResponse(response);
  }

  Future<RequestResponse<Map<String, dynamic>>> put([String? url]) async {
    _url = _buildUrl(url);

    var response = await http.put(
        Uri.parse(_url!),
        body: _body,
        headers: _headers);

    if (response.statusCode == 401 && !_hasRetried) {
      var refreshResponse = await _handleUnauthorizedResponse();
      if (refreshResponse.isSuccess) {
        return put(_url);
      }
      return RequestResponse.error("Failed to refresh token", 401);
    }

    return _getJsonFromResponse(response);
  }

  Future<RequestResponse<Map<String, dynamic>>> delete([String? url]) async {
    _url = _buildUrl(url);

    final response = await http.delete(Uri.parse(_url!), headers: _headers);

    if (response.statusCode == 401 && !_hasRetried) {
      var refreshResponse = await _handleUnauthorizedResponse();
      if (refreshResponse.isSuccess) {
        return delete(_url);
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
}