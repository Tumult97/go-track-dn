import 'package:flutter/material.dart';

import '../models/item.model.dart';

class ItemCard extends StatelessWidget {
  final Item item;

  const ItemCard({super.key, required this.item, required this.action});

  final VoidCallback action;
  
  @override
  Widget build(BuildContext context) {
    return Card(
      color: Colors.purple[50],
      child: Padding(
        padding: const EdgeInsets.all(20.0),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          spacing: 10.0,
          children: [
            Row(
              mainAxisAlignment: MainAxisAlignment.spaceBetween,
              children: [
                Text(
                  item.name,
                  style: TextStyle(
                    fontSize: 20,
                    fontWeight: FontWeight.bold,
                  ),
                ),
                IconButton.filledTonal(onPressed: action, icon: Icon(Icons.edit))
              ],
            ),
            Text(item.description),
            Row(
              mainAxisAlignment: MainAxisAlignment.spaceBetween,
              children: [
              Text('Value: ${item.price.toString()}'),
              Text('Quantity: ${item.quantity ?? ''}'),
            ],)
          ],
        ),
      ),
    );
  }
}
