'use strict';

/**
 * @ngdoc service
 * @name dropkickApp.products
 * @description
 * # products
 * Service in the dropkickApp.
 */
angular.module('dropkickApp')
  .service('products', function products($http) {
    // AngularJS will instantiate a singleton by calling "new" on this function
    return({
      	getProducts: getProducts
    });

	function getProducts(){
    	return getHttpGet(configuration.base_url + "/products");
    }

    
  });
