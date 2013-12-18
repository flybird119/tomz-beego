if (property.equals(RepeatPrintDialog.COLUMNPROPERTIES[1])) {
	return view.getRowNumber(); // 行数
} else if (property.equals(RepeatPrintDialog.COLUMNPROPERTIES[2])) {
	return view.getOrderContent(); // 医嘱内容
} else if (property.equals(RepeatPrintDialog.COLUMNPROPERTIES[3])) {
	return view.getRoaContent(); // 医嘱用法
} else if (property.equals(RepeatPrintDialog.COLUMNPROPERTIES[4])) {
	return view.getGroupedString(); // 组套医嘱打印字符
} else if (property.equals(RepeatPrintDialog.COLUMNPROPERTIES[5])) {
	return StringUtil.getInstance().timestamp2String(view.getStartedDateTime()); // 医嘱开立时间
} else if (property.equals(RepeatPrintDialog.COLUMNPROPERTIES[6])) {
	return view.getStartedSign(); // 医嘱开立签字
} else if (property.equals(RepeatPrintDialog.COLUMNPROPERTIES[7])) {
	return StringUtil.getInstance().timestamp2String(view.getAuditedDateTime()); // 医嘱审核时间
} else if (property.equals(RepeatPrintDialog.COLUMNPROPERTIES[8])) {
	return view.getAuditedSign(); // 医嘱审核签字
} else if (property.equals(RepeatPrintDialog.COLUMNPROPERTIES[9])) {
	return StringUtil.getInstance().timestamp2String(view.getStoppedDateTime()); // 医嘱停止时间
} else if (property.equals(RepeatPrintDialog.COLUMNPROPERTIES[10])) {
	return view.getStoppedSign(); // 医嘱停止签字
} else if (property.equals(RepeatPrintDialog.COLUMNPROPERTIES[11])) {
	return StringUtil.getInstance().timestamp2String(view.getStoppedAuditedDateTime()); // 医嘱停止审核时间
} else if (property.equals(RepeatPrintDialog.COLUMNPROPERTIES[12])) {
	return view.getStoppedAuditedSign(); // 医嘱停止审核签字
} else if (property.equals(RepeatPrintDialog.COLUMNPROPERTIES[13])) {
	return view.getCancelledStatus(); // 停止/作废标识
}

TableItem tableItem = (TableItem) element;
Noday noday = (Noday) tableItem.getData();
if (property.equals(COLUMNPROPERTIES[0])) {
    noday.setName((String)value);
} else if (property.equals(COLUMNPROPERTIES[1])) {
    noday.setCount(Integer.parseInt(value.toString()));
} else if (property.equals(COLUMNPROPERTIES[2])) {
    noday.setOpen((Boolean)value);
} else if (property.equals(COLUMNPROPERTIES[3])) {
    noday.setType((Integer)value);
} else if (property.equals(COLUMNPROPERTIES[4])) {
    noday.setData1((String)value);
} else if (property.equals(COLUMNPROPERTIES[5])) {
    noday.setData2((RGB)value);
} else if (property.equals(COLUMNPROPERTIES[7])) {
    noday.setData4((String)value);
}
viewer.refresh();