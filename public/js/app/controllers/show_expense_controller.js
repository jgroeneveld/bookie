(function() {
  var app = angular.module("bookie");

  app.controller("ShowExpenseController", function($scope, $routeParams, $location, ApiClient) {
    var expenseId = $routeParams["id"];

    ApiClient.getExpense(expenseId).then(function(expense) {
      $scope.expense = expense;
    });

    $scope.ConfirmDelete = function() {
      alert("Theoretisch jetzt geloescht :D " + expenseId);
    };
  });

})();
