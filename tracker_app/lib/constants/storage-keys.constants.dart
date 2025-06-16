import 'package:flutter_secure_storage/flutter_secure_storage.dart';

class StorageKeys {
  static String authTokenKey = "AUTH_TOKEN";
  static String refreshTokenKey = "AUTH_TOKEN";

  static AndroidOptions androidOptions = const AndroidOptions(
    encryptedSharedPreferences: true,
  );
}