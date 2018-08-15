//SPDX-License-Identifier: Apache-2.0

var tuna = require('./controller.js');

module.exports = function(app){

  app.get('/get_product/:id', function(req, res){
    tuna.get_product(req, res);
  });
  app.get('/get_all_product/:product_type', function(req, res){
    tuna.get_all_product_by_type(req, res);
  });
  app.get('/add_product/:product', function(req, res){
    tuna.add_product(req, res);
  });
  app.get('/get_all_product', function(req, res){
    tuna.get_all_product(req, res);
  });
  app.get('/change_holder/:holder', function(req, res){
    tuna.change_holder(req, res);
  });
}
