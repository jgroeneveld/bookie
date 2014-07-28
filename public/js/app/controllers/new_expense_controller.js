(function() {
  var app = angular.module("bookie");

  app.controller("NewExpenseController", function($location, $scope, ApiClient) {
    $scope.expense = {
      Date: dateToYMD(new Date(Date.now())),
      Amount: "",
      Category: "Edeka",
    };

    $scope.Categories = ["Edeka", "Lidl", "Bäcker", "Sonstiges", "Auto", "Penny", "Drogerie", "Markt"];

    $scope.Submit = function(user) {
      $scope.expense.User = user;
      ApiClient.createExpense($scope.expense).then(function() {
        $location.path("/");
      });
    };
  });

  function dateToYMD(date) {
    var d = date.getDate();
    var m = date.getMonth() + 1;
    var y = date.getFullYear();
    return '' + y + '-' + (m<=9 ? '0' + m : m) + '-' + (d <= 9 ? '0' + d : d);
  }

})();
