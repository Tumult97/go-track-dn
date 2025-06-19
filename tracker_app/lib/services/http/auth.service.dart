import 'package:flutter_secure_storage/flutter_secure_storage.dart';
import 'package:jwt_decoder/jwt_decoder.dart';

import '../../constants/storage-keys.constants.dart';
import '../../models/auth-response.model.dart';
import '../../models/register-response.model.dart';
import 'http-client-builder.dart';

class AuthService {
  static final _controller = "auth";
  static final FlutterSecureStorage _storage = FlutterSecureStorage();

  static Future<bool> isAuthenticated() async {
    try {
      var authToken = await _storage.read(key: StorageKeys.authTokenKey);

      if (authToken == null) {
        return false;
      }

      if (JwtDecoder.isExpired(authToken)) {
        await _storage.delete(key: StorageKeys.authTokenKey);
        return false;
      }

      return true;
    } catch (e) {
      return false;
    }
  }

  static Future<AuthResponse> login(String userName, String password) async {
    try {
      Map<String, String> loginRequest = {
        'email': userName,
        'password': password,
      };

      var jsonResponse = await HttpClientBuilder()
          .withUrl("$_controller/login")
          .isAnonymous()
          .withBody(loginRequest)
          .post();

      if (!jsonResponse.isSuccess) {
        throw Exception(jsonResponse.message);
      }

      var response = AuthResponse.fromJson(jsonResponse.data!);

      try {
        await _storage.delete(key: StorageKeys.authTokenKey);
        await _storage.write(key: StorageKeys.authTokenKey, value: response.accessToken);

        await _storage.delete(key: StorageKeys.refreshTokenKey);
        await _storage.write(key: StorageKeys.refreshTokenKey, value: response.refreshToken);
      } catch (storageError) {
        throw Exception('Failed to save authentication tokens: $storageError');
      }

      return response;
    } catch(e) {
      rethrow;
    }
  }

  static Future<bool> refresh() async {
    var refreshToken = await _storage.read(key: StorageKeys.refreshTokenKey);

    if (refreshToken == null) {
      return false;
    }

    var request = {
      'refresh_token': refreshToken
    };

    var response = await HttpClientBuilder()
        .withUrl("$_controller/refresh")
        .isAnonymous()
        .withBody(request)
        .post();

    if(!response.isSuccess){
      return false;
    }

    var responseData = AuthResponse.fromJson(response.data!);

    try {
      await _storage.delete(key: StorageKeys.authTokenKey);
      await _storage.write(key: StorageKeys.authTokenKey, value: responseData.accessToken);

      await _storage.delete(key: StorageKeys.refreshTokenKey);
      await _storage.write(key: StorageKeys.refreshTokenKey, value: responseData.refreshToken);
    } catch (storageError) {
      return false;
    }

    return true;
  }

  static Future<RegisterResponse> register(String firstName, String lastName, String email, String password) async {
    var request = {
      'first_name': firstName,
      'last_name': lastName,
      'email': email,
      'password': password,
    };

    var jsonResponse = await HttpClientBuilder()
        .withUrl('$_controller/register')
        .withBody(request)
        .isAnonymous()
        .post();

    if (!jsonResponse.isSuccess) {
      throw Exception(jsonResponse.message);
    }

    var response = RegisterResponse.fromJson(jsonResponse.data!);

    return response;
  }
}