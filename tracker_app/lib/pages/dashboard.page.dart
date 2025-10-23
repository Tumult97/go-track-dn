import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:tracker_app/router/routes.constants.dart';

class Dashboard extends StatefulWidget {
  final Widget child; // 👈 nested page content

  const Dashboard({super.key, required this.child});

  @override
  State<Dashboard> createState() => _DashboardState();
}

class _DashboardState extends State<Dashboard> {
  int _selectedIndex = 0;

  @override
  Widget build(BuildContext context) {
    final location = GoRouterState.of(context).uri.toString();
    _selectedIndex = _getSelectedIndex(location);

    return Scaffold(
      appBar: AppBar(
        title: const Text('Dashboard'),
        centerTitle: true,
      ),
      bottomNavigationBar: NavigationBar(
        selectedIndex: _selectedIndex,
        onDestinationSelected: (index) {
          setState(() => _selectedIndex = index);
          if (index == 0) context.go(Routes.items);
          if (index == 1) context.go(Routes.locations);
        },
        destinations: const [
          NavigationDestination(
            icon: Icon(Icons.inventory_2),
            label: 'Items',
          ),
          NavigationDestination(
            icon: Icon(Icons.location_on),
            label: 'Locations',
          ),
        ],
      ),
      body: widget.child,
    );
  }

  int _getSelectedIndex(String location) {
    if (location.contains('/locations')) return 1;
    return 0;
  }
}
