(function() {

  var app = angular.module("bookie", ['ngRoute', 'ngAnimate'])

  app.config(function($routeProvider) {

    $routeProvider.when("/expenses/new", {
      templateUrl: '/templates/new_expense.html',
      controller: 'NewExpenseController',
    });

    $routeProvider.otherwise({
      templateUrl: '/templates/expenses.html',
      controller: 'ExpensesController',
    });

  });

})();
