import 'package:tracker_app/models/location.model.dart';
import 'http/http-client-builder.dart';

class LocationService {
  static final _controller = 'location';

  /// Get all locations
  Future<List<Location>> $getLocations() async {
    final jsonResponse = await HttpClientBuilder().get(_controller);

    if (!jsonResponse.isSuccess) {
      throw Exception(jsonResponse.message);
    }

    final dataList = jsonResponse.data?['data-list'] as List<dynamic>?;

    if (dataList == null) return [];

    return dataList.map((location) => Location.fromJson(location)).toList();
  }

  /// Get location by ID
  Future<Location> $getLocationById(int id) async {
    final jsonResponse = await HttpClientBuilder().get('$_controller/$id');

    if (!jsonResponse.isSuccess) {
      throw Exception(jsonResponse.message);
    }

    if (jsonResponse.data == null) {
      throw Exception('Location not found');
    }

    return Location.fromJson(jsonResponse.data!);
  }

  /// Create a new location
  Future<Location> $createLocation(Location location) async {
    final response =
    await HttpClientBuilder().withBody(location.toJson()).post(_controller);

    if (!response.isSuccess) {
      throw Exception(response.message);
    }

    if (response.data == null) {
      throw Exception('No data returned from server');
    }

    return Location.fromJson(response.data!);
  }

  /// Update an existing location
  Future<Location> $updateLocation(Location location) async {
    final response =
    await HttpClientBuilder().withBody(location.toJson()).put(_controller);

    if (!response.isSuccess) {
      throw Exception(response.message);
    }

    if (response.data == null) {
      throw Exception('No data returned from server');
    }

    return Location.fromJson(response.data!);
  }

  /// Delete a location by ID
  Future<void> $deleteLocation(int id) async {
    final response = await HttpClientBuilder().delete('$_controller/$id');

    if (!response.isSuccess) {
      throw Exception(response.message);
    }
  }
}
