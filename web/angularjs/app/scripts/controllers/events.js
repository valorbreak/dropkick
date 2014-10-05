'use strict';

/**
 * @ngdoc function
 * @name dropkickApp.controller:EventsCtrl
 * @description
 * # EventsCtrl
 * Controller of the dropkickApp
 */
angular.module('dropkickApp')
  .controller('EventsCtrl', function ($scope,$http,$route) {

  	// Send a GET request to a url
  	// assign the reponse body to $scope for the views to update

    $http({method: 'GET', url: 'http://www.reddit.com/r/DotA2.json'}).
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
