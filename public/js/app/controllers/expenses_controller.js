(function() {
  var app = angular.module("bookie");

  app.controller("ExpensesController", function($scope, $location, ApiClient) {
    ApiClient.getExpenses().then(function(expenses) {
      $scope.expenses = expenses;
    });

    $scope.showExpense = function(expense) {
      $location.path('/expenses/' + expense.ID);
    };
  });

})();
