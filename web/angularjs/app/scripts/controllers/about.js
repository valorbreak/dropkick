'use strict';

/**
 * @ngdoc function
 * @name dropkickApp.controller:AboutCtrl
 * @description
 * # AboutCtrl
 * Controller of the dropkickApp
 */
angular.module('dropkickApp')
  .controller('AboutCtrl', function ($scope,$http,$route) {
    $http({method: 'GET', url: 'http://muledev1.hhmi.org:7015/PS/people/12345'}).
      success(function(data, status, headers, config) {
      	$scope.awesomeThings = data;
      	if(data.data.children){
      		$scope.list = data.data.children;
      	}
        // this callback will be called asynchronously
        // when the response is available
      }).
    error(function(data, status, headers, config) {
      // called asynchronously if an error occurs

      // or server returns response with an error status.
  	});
  });
