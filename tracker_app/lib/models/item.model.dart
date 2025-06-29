class Item {
  final String name;
  final String description;
  final String? quantity;
  final double price;
  final bool isPerItem;
  final int? locationId;

  Item({
    required this.name,
    required this.description,
    this.quantity,
    required this.price,
    required this.isPerItem,
    this.locationId,
  });

  factory Item.fromJson(Map<String, dynamic> json) {
    return Item(
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
      'name': name,
      'description': description,
      'quantity': quantity,
      'price': price,
      'isPerItem': isPerItem,
      'locationId': locationId,
    };
  }
}
