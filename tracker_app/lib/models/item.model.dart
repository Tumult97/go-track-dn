import 'package:tracker_app/models/common/entity-base.model.dart';

class Item extends EntityBase {
  final String name;
  final String description;
  final String? quantity;
  final double price;
  final bool isPerItem;
  final int? locationId;
  final DateTime? boughtDate;

  Item({
    super.id,
    super.created,
    required this.name,
    required this.description,
    this.quantity,
    required this.price,
    required this.isPerItem,
    this.locationId,
    this.boughtDate
  });

  factory Item.fromJson(Map<String, dynamic> json) {
    return Item(
      id: json['id'] as int,
      name: json['name'] as String,
      description: json['description'] as String,
      quantity: json['quantity'] as String?,
      price: (json['price'] as num).toDouble(),
      isPerItem: json['isPerItem'] as bool,
      locationId: json['locationId'] as int?,
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'id': id,
      'created': created,
      'name': name,
      'description': description,
      'quantity': quantity,
      'price': price,
      'isPerItem': isPerItem,
      'locationId': locationId,
      'boughtDate': boughtDate
    };
  }
}
