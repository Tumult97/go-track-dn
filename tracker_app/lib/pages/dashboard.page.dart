import 'dart:convert';
import 'dart:core';

import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:tracker_app/router/routes.constants.dart';
import 'package:tracker_app/services/item.service.dart';

import '../components/item-card.dart';
import '../models/item.model.dart';

class Dashboard extends StatefulWidget {

  const Dashboard({super.key});

  @override
  State<Dashboard> createState() => _DashboardState();
}

class _DashboardState extends State<Dashboard> {
  List<Item> _items = [];

  final ItemService _itemService = ItemService();

  @override
  Widget build(BuildContext context) {

    String test = json.encode(_items).toString();

    return Scaffold(
      body: Padding(
        padding: const EdgeInsets.all(20.0),
        child: ListView(
          children: _buildItemCards(),
        ),
      ),
      floatingActionButton: FloatingActionButton(
        onPressed: () => _openItemEdit(RouteNames.itemAdd),
        child: const Icon(Icons.add),
      ),
    );
  }

  @override
  void initState() {
    super.initState();
    _loadData();
  }

  Future _loadData() async {
    var items = await _itemService.$getItems();
    setState(() {
      _items = items;
    });
  }

  List<ItemCard> _buildItemCards() {
    return _items.map((item) => ItemCard(
        item: item,
        action: () => _openItemEdit(RouteNames.itemEdit, item),
    )).toList();
  }

  void _openItemEdit(String routeName, [Item? item]) async {
    var response = await context.pushNamed<Item?>(routeName, extra: item);

    if (response != null) {
      _loadData();
    }
  }
}
