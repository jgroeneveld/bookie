(function() {

  var app = angular.module("bookie", ['ngRoute'])

  app.config(function($routeProvider) {

    $routeProvider.otherwise({
      templateUrl: '/templates/expenses.html',
      controller: 'ExpensesController',
    });

  });

})();
