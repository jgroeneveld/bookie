(function() {
  var app = angular.module("bookie");

  app.controller("ExpensesController", function($scope, $http) {
    $http.get("/api/expenses").then(function(response) {
      var expenses = response.data;
      $scope.expenses = expenses;
    });
  });

})();
