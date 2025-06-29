import 'package:flutter/material.dart';

import '../models/item.model.dart';

class ItemCard extends StatelessWidget {
  final Item item;

  const ItemCard({super.key, required this.item});
  
  @override
  Widget build(BuildContext context) {
    return Card(
      child: Center(
        child: Text(item.name),
      ),
    );
  }
}
