// SPDX-License-Identifier: Apache-2.0

'use strict';

var app = angular.module('application', []);

// Angular Controller
app.controller('appController', function($scope, appFactory){

	$("#success_holder").hide();
	$("#success_create").hide();
	$("#error_holder").hide();
	$("#error_query").hide();
	$("#error_consumer_query").hide();
	
	$scope.queryAllProduct = function(){

		appFactory.queryAllProduct(function(data){
			var array = [];
			for (var i = 0; i < data.length; i++){
				parseInt(data[i].Key);
				data[i].Record.Key = parseInt(data[i].Key);
				array.push(data[i].Record);
			}
			array.sort(function(a, b) {
			    return parseFloat(a.Key) - parseFloat(b.Key);
			});
			$scope.all_product = array;
		});
	}

	$scope.queryProduct = function(){

		var id = $scope.product_type;

		appFactory.queryProduct(id, function(data){
			$scope.query_product = data;
			$scope.query_product.Key = id;

			if ($scope.product == "Could not locate product"){
				console.log()
				$("#error_query").show();
			} else{
				$("#error_query").hide();
			}
		});
	}

	$scope.queryAllProductByType = function(){

		var product_type = $scope.consumer_product_type;

		appFactory.queryAllProductByType(product_type, function(data){
			var array = [];
			for (var i = 0; i < data.length; i++){
				parseInt(data[i].Key);
				data[i].Record.Key = parseInt(data[i].Key);
				array.push(data[i].Record);
			}
			array.sort(function(a, b) {
			    return parseFloat(a.Key) - parseFloat(b.Key);
			});
			$scope.query_products = array;
		});
	}

	$scope.consumerQueryAllProductByType = function(){

		var product_type = $scope.consumer_product_type;

		appFactory.queryAllProductByType(product_type, function(data){
			var array = [];
			for (var i = 0; i < data.length; i++){
				parseInt(data[i].Key);
				data[i].Record.Key = parseInt(data[i].Key);
				array.push(data[i].Record);
			}
			array.sort(function(a, b) {
			    return parseFloat(a.Key) - parseFloat(b.Key);
			});
			$scope.consumer_query_products = array;
		});
	}

	$scope.recordProduct= function(){

		appFactory.recordProduct($scope.product, function(data){
			$scope.create_product = data;
			$("#success_create").show();
		});
	}

	$scope.changeHolder = function(){

		appFactory.changeHolder($scope.holder, function(data){
			$scope.change_holder = data;
			if ($scope.change_holder == "Error: no product catch found"){
				$("#error_holder").show();
				$("#success_holder").hide();
			} else{
				$("#success_holder").show();
				$("#error_holder").hide();
			}
		});
	}

	$scope.checkAll = function() {
		angular.forEach($scope.consumer_query_products, function(product) {
		  product.select = $scope.selectAll;
		});
	};

});

// Angular Factory
app.factory('appFactory', function($http){
	
	var factory = {};

    factory.queryAllProductByType = function(product_type,callback){

    	$http.get('/get_all_product/' + product_type).success(function(output){
			callback(output)
		});
	}

	factory.queryAllProduct = function(callback){

    	$http.get('/get_all_product/').success(function(output){
			callback(output)
		});
	}

	factory.queryProduct = function(id, callback){
    	$http.get('/get_product/'+id).success(function(output){
			callback(output)
		});
	}

	factory.recordProduct = function(data, callback){

		var product = data.id + "-" + data.producing_area + "-" + data.product_type + "-" + data.product_unit + "-" + data.unit_price + "-" + data.quantity + "-" + data.timestamp + "-" + data.holder;

    	$http.get('/add_product/'+product).success(function(output){
			callback(output)
		});
	}

	factory.changeHolder = function(data, callback){

		var holder = data.id + "-" + data.name;

    	$http.get('/change_holder/'+holder).success(function(output){
			callback(output)
		});
	}

	return factory;
});


