(function() {
  var app = angular.module("bookie");

  function spendersFromReport(report) {
    var spenders = []
    for (var username in report.AmountByUsers) {
      spenders.push({
        User: username,
        Amount: report.AmountByUsers[username],
      });
    }
    return spenders;
  }

  function sortSpendersByAmount(spenders) {
    spenders.sort(function(a,b) {
      return (b.Amount - a.Amount);
    });
  }

  app.controller("TotalAmountController", function($scope, ApiClient) {
    ApiClient.getReport().then(function(report) {
      var spenders = spendersFromReport(report);
      sortSpendersByAmount(spenders)
      $scope.biggestSpender = spenders[0];
      $scope.lesserSpender = spenders[spenders.length-1];
    });
  });

})();
