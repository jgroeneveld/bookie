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

      createExpense: function(expense) {
        return $http.post("/api/expenses").then(function(response) {
          console.log(response);
        });
      },

    };
  });
})();
