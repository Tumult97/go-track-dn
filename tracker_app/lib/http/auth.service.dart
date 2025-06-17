import 'package:flutter_secure_storage/flutter_secure_storage.dart';
import 'package:jwt_decoder/jwt_decoder.dart';
import 'package:tracker_app/constants/storage-keys.constants.dart';
import 'package:tracker_app/http/http-client-builder.dart';
import 'package:tracker_app/models/auth-response.model.dart';

class AuthService {
  static final _controller = "auth";
  static final FlutterSecureStorage _storage = FlutterSecureStorage();

  static Future<bool> isAuthenticated() async {
    try {
      var authToken = await _storage.read(key: StorageKeys.authTokenKey);

      if (authToken == null) {
        return false;
      }

      return !JwtDecoder.isExpired(authToken);
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

      var response = AuthResponse.fromJson(jsonResponse);

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

    var jsonResponse = await HttpClientBuilder()
        .withUrl("$_controller/refresh")
        .isAnonymous()
        .withBody(request)
        .post();

    var response = AuthResponse.fromJson(jsonResponse);

    try {
      await _storage.delete(key: StorageKeys.authTokenKey);
      await _storage.write(key: StorageKeys.authTokenKey, value: response.accessToken);

      await _storage.delete(key: StorageKeys.refreshTokenKey);
      await _storage.write(key: StorageKeys.refreshTokenKey, value: response.refreshToken);
    } catch (storageError) {
      return false;
    }

    return true;
  }
}