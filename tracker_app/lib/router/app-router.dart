import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:tracker_app/pages/auth/login.page.dart';
import 'package:tracker_app/pages/auth/register.page.dart';
import 'package:tracker_app/pages/dashboard.page.dart';
import 'package:tracker_app/pages/items/item-edit.page.dart';
import 'package:tracker_app/pages/items/items-list.dart';
import 'package:tracker_app/pages/locations/location-list.dart';
import 'package:tracker_app/router/routes.constants.dart';
import '../models/item.model.dart';
import '../services/http/auth.service.dart';

class AppRouter {
  static final router = GoRouter(
    initialLocation: Routes.items,
    redirect: (context, state) async {
      final isAuthenticated = await AuthService.isAuthenticated();
      final isLoggingIn = state.matchedLocation == Routes.signIn || state.matchedLocation == Routes.register;

      if (!isAuthenticated && !isLoggingIn) {
        return Routes.signIn;
      }

      if (isAuthenticated && isLoggingIn) {
        return Routes.items;
      }

      return null;
    },
    routes: [
      GoRoute(
        path: Routes.home,
        name: RouteNames.home,
        redirect: (context, state) => Routes.items,
      ),
      ShellRoute(
        builder: (context, state, child) {
          return Dashboard(child: child);
        },
        routes: [
          GoRoute(
            path: Routes.items,
            name: RouteNames.items,
            builder: (context, state) => ItemList(),
          ),
          GoRoute(
            path: Routes.locations,
            name: RouteNames.locations,
            builder: (context, state) => LocationList(),
          ),
        ],
      ),

      GoRoute(
        name: RouteNames.signIn,
        path: Routes.signIn,
        builder: (context, state) => LoginPage(),
      ),
      GoRoute(
        name: RouteNames.register,
        path: Routes.register,
        builder: (context, state) => RegisterPage(),
      ),

      GoRoute(
        name: RouteNames.itemAdd,
        path: Routes.addItem,
        builder: (context, state) => ItemEdit(),
      ),
      GoRoute(
        name: RouteNames.itemEdit,
        path: Routes.editItem,
        builder: (context, state) {
          final item = state.extra as Item;
          return ItemEdit(item: item);
        },
      ),
      GoRoute(
        name: RouteNames.itemView,
        path: Routes.viewItem,
        builder: (context, state) => const Placeholder(),
      ),
    ],
  );
}