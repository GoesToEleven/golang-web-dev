var PODetail = function(obj){
  if (!obj) {
    obj = {}
  }
  this.id = obj.id || 0;
  this.line = obj.line || 0;
  this.partNumber = obj.partNumber || "";
  this.description = obj.description || "";
  this.quantity = obj.quantity || 0;
  this.dueDate = obj.dueDate || null;
  this.unitOfMeasure = obj.unitOfMeasure || "";
  this.pricePer = obj.pricePer || 0;
  this.status = obj.status || "";
}
