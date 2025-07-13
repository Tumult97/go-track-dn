import 'package:flutter/cupertino.dart';
import 'package:go_router/go_router.dart';
import 'package:tracker_app/pages/auth/login.page.dart';
import 'package:tracker_app/pages/dashboard.page.dart';
import 'package:tracker_app/pages/items/item-edit.page.dart';
import 'package:tracker_app/router/routes.constants.dart';
import '../models/item.model.dart';
import '../pages/auth/register.page.dart';
import '../services/http/auth.service.dart';

class AppRouter {
  static final router = GoRouter(
    initialLocation: Routes.defaultRoute,
    redirect: (context, state) async {
      final isAuthenticated = await AuthService.isAuthenticated();

      final isLoggingIn = state.matchedLocation == Routes.signIn || state.matchedLocation == Routes.register;

      if (!isAuthenticated && !isLoggingIn) {
        return Routes.signIn;
      }

      if (isAuthenticated && isLoggingIn) {
        return Routes.defaultRoute;
      }

      return null;
    },
    routes: [
      GoRoute(
          path: Routes.defaultRoute,
          builder: (context, state) => Dashboard(),
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
        name: RouteNames.home,
        path: Routes.home,
        builder: (context, state) => Dashboard(),
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
        builder: (context, state) => Placeholder(),
      )
    ]
  );
}