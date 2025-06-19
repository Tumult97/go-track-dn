import 'package:flutter_secure_storage/flutter_secure_storage.dart';

import '../constants/storage-keys.constants.dart';

class LocalStorageService {
  final FlutterSecureStorage _storage = FlutterSecureStorage();

  Future saveValue(String key, String value) {
    _storage.delete(key: key);
    return _storage.write(key: key, value: value);
  }

  Future<String?> fetchValue<T>(String value) {
    return _storage.read(key: value);
  }

  Future deleteValue(String key) {
    return _storage.delete(key: key);
  }

  Future logOut() async {
    await _storage.delete(key: StorageKeys.authTokenKey);
    await _storage.delete(key: StorageKeys.refreshTokenKey);
  }
}