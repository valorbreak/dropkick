'use strict';

/**
 * @ngdoc function
 * @name dropkickApp.controller:RedditdotaCtrl
 * @description
 * # RedditdotaCtrl
 * Controller of the dropkickApp
 */
var controllers = {};

controllers = { 
	RedditdotaCtrl : function ($scope, redditDota) {
    		redditDota.getLatestReddit().then( function(response){
    		$scope.list = response.data.children;
    		$scope.name = 'ctrl1';	
    	});
	},
	RedditdotaCtrl2 : function ($scope, redditDota) {
    		redditDota.getLatestReddit().then( function(response){
    		$scope.list = response.data.children;
    		$scope.name = 'ctrl1';	
    	});
	}
}

/*
controllers.RedditdotaCtrl = function ($scope, redditDota) {
		redditDota.getLatestReddit().then( function(response){
		$scope.list = response.data.children;
		$scope.name = 'ctrl1';	
    });
}
*/


var dropkickApp = angular.module('dropkickApp');
dropkickApp.controller(controllers);

console.log(app,'app');
console.log(dropkickApp,'dk');
console.log(dependencies,'dk');