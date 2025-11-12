import 'package:tracker_app/models/common/entity-base.model.dart';

class Location extends EntityBase {
  final String name;
  final String description;
  final String addressHome;
  final String addressStreet;
  final String addressSuburb;
  final String addressCity;

  Location({
    super.id,
    super.created,
    required this.name,
    required this.description,
    required this.addressHome,
    required this.addressStreet,
    required this.addressSuburb,
    required this.addressCity,
  });

  factory Location.fromJson(Map<String, dynamic> json) {
    return Location(
      id: json['id'] ?? 0,
      name: json['name'] ?? '',
      description: json['description'] ?? '',
      addressHome: json['addressHome'] ?? '',
      addressStreet: json['addressStreet'] ?? '',
      addressSuburb: json['addressSuburb'] ?? '',
      addressCity: json['addressCity'] ?? '',
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'id': id,
      'name': name,
      'description': description,
      'addressHome': addressHome,
      'addressStreet': addressStreet,
      'addressSuburb': addressSuburb,
      'addressCity': addressCity,
    };
  }
}
