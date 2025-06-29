import 'package:http/http.dart' as http;
import 'package:tracker_app/enums/method-type.dart';
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
}