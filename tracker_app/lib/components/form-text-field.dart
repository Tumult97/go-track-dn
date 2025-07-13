import 'package:flutter/material.dart';

class FormTextField extends StatelessWidget {
  const FormTextField({super.key, required this.controller, this.validator, this.maxLines = 5});

  final TextEditingController controller;
  final int maxLines;
  final String? Function(dynamic)? validator;

  @override
  Widget build(BuildContext context) {
    return TextFormField(
      controller: controller,
      decoration: InputDecoration(
          labelText: 'Description',
          border: OutlineInputBorder(
            borderRadius: BorderRadius.circular(10.0)
          )
      ),
      maxLines: maxLines,
      validator: validator,
    );
  }
}
