(function() {
  var app = angular.module("bookie");

  app.controller("ExpensesController", function($scope, ApiClient) {
    ApiClient.getExpenses().then(function(expenses) {
      $scope.expenses = expenses;
    });
  });

})();
