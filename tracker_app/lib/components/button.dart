import 'package:flutter/material.dart';

class Button extends StatelessWidget {
  const Button({super.key, required this.text, required this.action});

  final String text;
  final VoidCallback? action;

  @override
  Widget build(BuildContext context) {
    return TextButton(
      onPressed: action,
      style: ElevatedButton.styleFrom(
          backgroundColor: Colors.deepPurple,
          foregroundColor: Colors.white,
      ),
      child: Text(text),
    );
  }
}
