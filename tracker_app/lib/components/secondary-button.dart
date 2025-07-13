import 'package:flutter/material.dart';
import 'package:tracker_app/components/button.dart';

class SecondaryButton extends Button {
  const SecondaryButton({super.key, required super.text, required super.action});

  @override
  Widget build(BuildContext context) {
    return OutlinedButton(
      onPressed: action,
      style: OutlinedButton.styleFrom(
        shape: RoundedRectangleBorder(
          borderRadius: BorderRadius.circular(10.0)
        ),
      ),
      child: Text(text),
    );
  }
}
