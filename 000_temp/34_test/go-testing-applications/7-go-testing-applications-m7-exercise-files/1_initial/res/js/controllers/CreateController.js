(function() {
  var CreateController = function(ShippingService, CurrencyService, PurchaseOrderService, $scope, $state) {
    var self = this;

    ShippingService.getVendors().then(function(vendors) {
      self.vendors = vendors;
    });
    this.currentVendor = null;
    CurrencyService.getCurrencies().then(function(currencies) {
      self.currencies = currencies;
    });
    this.currency = null;
    ShippingService.getReceivers().then(function(receivers) {
      self.receivers = receivers;
    });
    this.currentReceiver = null;
    this.isDatePickerOpen = false;

    this.details = [];

    this.totalCost = 0;

    function updateTotalCost() {
      self.totalCost = self.details.reduce(function(value, detail) {
        return value + detail.pricePer * detail.quantity;
      }, 0);
    }

    this.toggleDatePicker = function(e) {
      this.isDatePickerOpen = !this.isDatePickerOpen;
    }

    this.addDetail = function() {
      var dtl = new PODetail();
      dtl.unwatcher = $scope.$watch(function() {return dtl;}, function() {
        updateTotalCost();
      }, true);
      this.details.push(dtl);
    }

    this.removeDetail = function(detail) {
      for (var i = 0; i < this.details.length; i++) {
        if (detail == this.details[i]) {
          this.details.splice(i,1);
          detail.unwatcher();
          break;
        }
      }
    }

    this.createOrder = function() {
      var po = new PurchaseOrder();

      po.receiver = this.currentReceiver;
      po.vendor = this.currentVendor;
      po.currency = this.currency;
      po.purpose = this.purpose;
      po.dueDate = this.dueDate;
      po.details = this.details.slice();

      PurchaseOrderService.createOrder(po).then(function(order) {
        self.clearForm();
        location.hash = "#/purchaseOrders/" + order.id + "/manage";
      })

    }

    this.clearForm = function() {
      this.currentReceiver = null;
      this.currentVendor = null;
      this.currency = null;
      this.purpose = "";
      this.dueDate = null;
      while(this.details.length) {
        this.removeDetail(this.details[0]);
      }
    }

    $scope.$watchCollection(function() {
      return self.details;
    }, function() {
      updateTotalCost();
    });

  }

  angular.module("app")
    .controller("CreateController", CreateController);
})();
