import 'package:tracker_app/models/item.model.dart';

import 'http/http-client-builder.dart';

class ItemService {
  static final _controller = 'items';

  Future<List<Item>> $getItems() async {
    var jsonResponse = await HttpClientBuilder()
        .get(_controller);

    if (!jsonResponse.isSuccess) {
      throw Exception(jsonResponse.message);
    }

    var dataItems = jsonResponse.data?['data-list'] as List<dynamic>;

    return dataItems.map((item) => Item.fromJson(item)).toList();
  }

  Future<Item> $createItem(Item item) async {
    var response = await HttpClientBuilder()
        .withBody(item.toJson())
        .post(_controller);

    if (!response.isSuccess) {
      throw Exception(response.message);
    }

    if (response.data == null) {
      throw Exception('No data returned from server');
    }

    return Item.fromJson(response.data!);
  }

  Future $updateItem(Item item) async {
    var response = await HttpClientBuilder()
        .withBody(item.toJson())
        .put(_controller);

    if (!response.isSuccess) {
      throw Exception(response.message);
    }

    if (response.data == null) {
      throw Exception('No data returned from server');
    }

    return Item.fromJson(response.data!);
  }
}