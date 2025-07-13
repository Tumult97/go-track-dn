import 'package:flutter/material.dart';
import 'package:tracker_app/components/button.dart';

class PrimaryButton extends Button {
  const PrimaryButton({super.key, required super.text, required super.action});

  @override
  Widget build(BuildContext context) {
    return ElevatedButton(
      onPressed: action,
      style: ElevatedButton.styleFrom(
          backgroundColor: Colors.deepPurple,
          foregroundColor: Colors.white,
          shape: RoundedRectangleBorder(
              borderRadius: BorderRadius.circular(10.0)
          ),
      ),
      child: Text(text),
    );
  }
}
