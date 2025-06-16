import 'package:flutter_secure_storage/flutter_secure_storage.dart';
import 'package:tracker_app/constants/storage-keys.constants.dart';
import 'package:tracker_app/http/http-client-builder.dart';
import 'package:tracker_app/models/auth-response.model.dart';

class AuthService {
  static final _controller = "auth";
  static final FlutterSecureStorage _storage = const FlutterSecureStorage();

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

      _storage.write(key: StorageKeys.authTokenKey, value: response.accessToken);
      _storage.write(key: StorageKeys.refreshTokenKey, value: response.refreshToken);

      return response;
    } catch(e) {
      rethrow;
    }
  }
}