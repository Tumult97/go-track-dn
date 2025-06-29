import 'package:flutter_secure_storage/flutter_secure_storage.dart';

class StorageKeys {
  static String authTokenKey = "AUTH_TOKEN";
  static String refreshTokenKey = "REFRESH_TOKEN";

  static AndroidOptions androidOptions = AndroidOptions(
    encryptedSharedPreferences: true,
  );
}