var PurchaseOrder = function(obj) {
  if (!obj) {
    obj = {}
  }
  this.id = obj.id || 0;
  this.receiver = obj.receiver || null;
  this.vendor = obj.vendor || null;
  this.currency = obj.currency || null;
  this.purpose = obj.purpose || "";
  this.dueDate = obj.dueDate || null;
  this.status = obj.status || null;
  this.details = obj.details || [];

}

PurchaseOrder.prototype.totalCost = function() {
  var result = 0;
  if (this.details) {
    result = this.details.reduce(function(v, detail) {
      return v + detail.quantity * detail.pricePer
    }, result);
  }
  return result;
}
