import 'package:flutter/material.dart';
import 'package:flutter/services.dart';

class FormTextBox extends StatelessWidget {
  const FormTextBox({super.key, required this.controller, this.label = 'Label', this.keyboardType = TextInputType.text, this.formatter, this.validator});

  final TextEditingController controller;
  final String? label;
  final TextInputFormatter? formatter;
  final TextInputType keyboardType;
  final String? Function(dynamic)? validator;

  @override
  Widget build(BuildContext context) {
    return TextFormField(
      controller: controller,
      decoration: InputDecoration(
        labelText: label,
        border: OutlineInputBorder(
          borderRadius: BorderRadius.circular(10.0),
        ),
      ),
      keyboardType: TextInputType.text,
      inputFormatters: formatter != null ? [formatter!] : null,
      validator: validator,
    );
  }
}
