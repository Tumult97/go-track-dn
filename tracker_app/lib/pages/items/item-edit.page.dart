import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:tracker_app/components/form-text-box.dart';
import 'package:tracker_app/components/form-text-field.dart';
import 'package:tracker_app/components/primary-button.dart';
import 'package:tracker_app/components/secondary-button.dart';
import 'package:tracker_app/constants/form-validators.dart';

import '../../models/item.model.dart';
import '../../services/item.service.dart';

class ItemEdit extends StatefulWidget {
  const ItemEdit({super.key, this.item});
  final Item? item;

  @override
  State<ItemEdit> createState() => _ItemEditState();
}

class _ItemEditState extends State<ItemEdit> {
  final TextEditingController _nameController = TextEditingController();
  final TextEditingController _descriptionController = TextEditingController();
  final TextEditingController _priceController = TextEditingController();
  final TextEditingController _quantityController = TextEditingController();

  bool isPerItem = false;
  bool isLoading = false;

  final _formKey = GlobalKey<FormState>();

  final _itemService = ItemService();

  @override
  Widget build(BuildContext context) {
    _setInputs();

    return Scaffold(
      appBar: AppBar(
        title: Text(widget.item == null ? 'Add Item' : 'Edit Item'),
      ),
      body: Form(
        key: _formKey,
        child: Padding(
          padding: const EdgeInsets.all(20.0),
          child: Column(
            mainAxisAlignment: MainAxisAlignment.spaceBetween,
            children: [
              Column(
                spacing: 10.0,
                children: [
                  FormTextBox(
                    controller: _nameController,
                    label: 'Name',
                    validator: FormValidators.notEmptyValidator,
                  ),
                  FormTextField(
                    controller: _descriptionController,
                    validator: FormValidators.notEmptyValidator,
                  ),
                  Row(
                    crossAxisAlignment: CrossAxisAlignment.start,
                    spacing: 10.0,
                    children: [
                      Expanded(
                        child: FormTextBox(
                          controller: _priceController,
                          label: 'Price',
                          validator: FormValidators.nonEmptyCurrencyValidator,
                          keyboardType: TextInputType.numberWithOptions(decimal: true),
                        ),
                      ),
                      Expanded(
                        child: FormTextBox(
                          controller: _quantityController,
                          label: 'Quantity',
                          validator: FormValidators.notEmptyValidator,
                          keyboardType: TextInputType.number,
                        ),
                      ),
                    ]
                  ),
                  InkWell(
                    onTap: () => setState(() => isPerItem = !isPerItem),
                    child: Row(
                      children: [
                        Checkbox(
                          value: isPerItem,
                          semanticLabel: 'Per item',
                          onChanged: (bool? value) =>
                              setState(() {
                                isPerItem = value!;
                              }),
                        ),
                        Text('Is Per Item')
                      ],
                    ),
                  ),
                ]
              ),
              Row(
                spacing: 10.0,
                children: [
                  Expanded(child: SecondaryButton(text: 'CANCEL', action: () => context.pop())),
                  Expanded(child: PrimaryButton(text: 'SAVE', action: _submitItem))
                ],
              ),
            ]
          ),
        ),
      ),
    );
  }

  void _setInputs(){
    if(widget.item == null){
      return;
    }

    _nameController.text = widget.item!.name;
    _descriptionController.text = widget.item!.description;
    _priceController.text = widget.item!.price.toString();
    _quantityController.text = widget.item!.quantity.toString();
  }

  Future _submitItem() async {
    var formValid = _formKey.currentState!.validate();
    if (!formValid) {
      return;
    }

    var item = _getItemFromForm();

    setState(() {
      isLoading = true;
    });

    var response = await (item.id == 0 ? _createItem(item) : _updateItem(item));

    _closePage(item);
  }

  void _closePage(Item? item) {
    context.pop(item);
  }

  Item _getItemFromForm(){
    return Item(
      id: widget.item?.id ?? 0,
      name: _nameController.text,
      description: _descriptionController.text,
      price: double.parse(_priceController.text),
      quantity: _quantityController.text,
      isPerItem: isPerItem,
    );
  }

  Future<Item?> _createItem(Item item) async {
    var response = await _itemService.$createItem(item);
    return response;
  }

  Future<Item?> _updateItem(Item item) async {
    var response = await _itemService.$createItem(item);
    return response;
  }
}
