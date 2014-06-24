(function() {
  var app = angular.module("bookie");

  app.factory("ApiClient", function($http) {
    return {

      getExpenses: function() {
        return $http.get("/api/expenses").then(function(response) {
          var expenses = response.data;
          return expenses;
        });
      },

      getReport: function() {
        return $http.get("/api/expenses/report").then(function(response) {
          var report = response.data;
          return report;
        });
      },

      createExpense: function(expense) {
        return $http({
          method: 'post',
          url: '/api/expenses',
          params: expense,
        }).then(function(response) {
          console.log(response);
        });
      },

    };
  });
})();
