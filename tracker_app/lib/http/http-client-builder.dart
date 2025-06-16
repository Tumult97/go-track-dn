import 'dart:convert';

import 'package:http/http.dart' as http;

class HttpClientBuilder {
  String? _url;
  Map<String, String> _queryParameters = {};
  bool _useAuth = true;
  Map<String, dynamic>? _body;

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

  Future<Map<String, dynamic>> get([String? url]) async {
    _url = _buildUrl(url);

    final response = await http.get(Uri.parse(_url!), headers: _headers);

    return _getJsonFromResponse(response);
  }

  Future<Map<String, dynamic>> post([String? url]) async {
    _url = _buildUrl(url);

    var response = await http.post(
        Uri.parse(_url!),
        body: jsonEncode(_body),
        headers: _headers);

    return _getJsonFromResponse(response);
  }

  Future<Map<String, dynamic>> put([String? url]) async {
    _url = _buildUrl(url);

    var response = await http.put(
        Uri.parse(_url!),
        body: _body,
        headers: _headers);

    return _getJsonFromResponse(response);
  }

  Future<Map<String, dynamic>> delete([String? url]) async {
    _url = _buildUrl(url);

    final response = await http.delete(Uri.parse(_url!), headers: _headers);

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

  Map<String, dynamic> _getJsonFromResponse(http.Response response) {
    if (response.statusCode == 200) {
      return json.decode(response.body);
    } else {
      throw Exception('Failed to load data: ${response.statusCode}');
    }
  }
}