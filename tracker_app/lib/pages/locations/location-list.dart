import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:tracker_app/models/location.model.dart';
import 'package:tracker_app/router/routes.constants.dart';
import 'package:tracker_app/services/location.service.dart';

class LocationList extends StatefulWidget {
  const LocationList({super.key});

  @override
  State<LocationList> createState() => _LocationListState();
}

class _LocationListState extends State<LocationList> {
  List<Location> _locations = [];
  final LocationService _locationService = LocationService();

  @override
  void initState() {
    super.initState();
    _loadData();
  }

  Future<void> _loadData() async {
    try {
      final locations = await _locationService.$getLocations();
      setState(() {
        _locations = locations;
      });
    } catch (e) {
      debugPrint('Error loading locations: $e');
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Padding(
        padding: const EdgeInsets.all(20.0),
        child: ListView(
          children: _buildLocationCards(),
        ),
      ),
      floatingActionButton: FloatingActionButton(
        onPressed: () {},
        child: const Icon(Icons.add),
      ),
    );
  }

  List<Widget> _buildLocationCards() {
    return _locations.map((location) {
      return Card(
        color: Colors.purple[50],
        child: Padding(
          padding: const EdgeInsets.all(20.0),
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            spacing: 10.0,
            children: [
              // Header: name + edit button
              Row(
                mainAxisAlignment: MainAxisAlignment.spaceBetween,
                children: [
                  Text(
                    location.name,
                    style: const TextStyle(
                      fontSize: 20,
                      fontWeight: FontWeight.bold,
                    ),
                  ),
                  IconButton.filledTonal(
                    onPressed: () {},
                    icon: const Icon(Icons.edit),
                  ),
                ],
              ),

              // Description
              Text(
                location.description.isNotEmpty
                    ? location.description
                    : 'No description provided.',
                style: const TextStyle(color: Colors.black87),
              ),

              const SizedBox(height: 10),

              // Address
              Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  Text(location.addressHome),
                  Text(location.addressStreet),
                  Text('${location.addressSuburb}, ${location.addressCity}'),
                ],
              ),
            ],
          ),
        ),
      );
    }).toList();
  }

  void _openLocationEdit(String routeName, [Location? location]) async {
    final response =
    await context.pushNamed<Location?>(routeName, extra: location);

    if (response != null) {
      _loadData();
    }
  }
}
