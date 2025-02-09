// Code generated by "stringer -linecomment -type ErrorCode"; DO NOT EDIT.

package handlererrors

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[errUnset-0]
	_ = x[errInternalError-1]
	_ = x[ErrBadValue-2]
	_ = x[ErrFailedToParse-9]
	_ = x[ErrUserNotFound-11]
	_ = x[ErrUnauthorized-13]
	_ = x[ErrTypeMismatch-14]
	_ = x[ErrAuthenticationFailed-18]
	_ = x[ErrIllegalOperation-20]
	_ = x[ErrNamespaceNotFound-26]
	_ = x[ErrIndexNotFound-27]
	_ = x[ErrUnsuitableValueType-28]
	_ = x[ErrConflictingUpdateOperators-40]
	_ = x[ErrCursorNotFound-43]
	_ = x[ErrNamespaceExists-48]
	_ = x[ErrDollarPrefixedFieldName-52]
	_ = x[ErrInvalidID-53]
	_ = x[ErrEmptyName-56]
	_ = x[ErrCommandNotFound-59]
	_ = x[ErrImmutableField-66]
	_ = x[ErrCannotCreateIndex-67]
	_ = x[ErrIndexAlreadyExists-68]
	_ = x[ErrInvalidOptions-72]
	_ = x[ErrInvalidNamespace-73]
	_ = x[ErrIndexOptionsConflict-85]
	_ = x[ErrIndexKeySpecsConflict-86]
	_ = x[ErrOperationFailed-96]
	_ = x[ErrDocumentValidationFailure-121]
	_ = x[ErrInvalidIndexSpecificationOption-197]
	_ = x[ErrInvalidPipelineOperator-168]
	_ = x[ErrClientMetadataCannotBeMutated-186]
	_ = x[ErrNotImplemented-238]
	_ = x[ErrIndexesWrongType-10065]
	_ = x[ErrDuplicateKeyInsert-11000]
	_ = x[ErrSetBadExpression-40272]
	_ = x[ErrStageGroupInvalidFields-15947]
	_ = x[ErrStageGroupID-15948]
	_ = x[ErrStageGroupMissingID-15955]
	_ = x[ErrStageLimitZero-15958]
	_ = x[ErrMatchBadExpression-15959]
	_ = x[ErrProjectBadExpression-15969]
	_ = x[ErrSortBadExpression-15973]
	_ = x[ErrSortBadValue-15974]
	_ = x[ErrSortBadOrder-15975]
	_ = x[ErrSortMissingKey-15976]
	_ = x[ErrGroupDuplicateFieldName-16406]
	_ = x[ErrStageUnwindWrongType-15981]
	_ = x[ErrExpressionWrongLenOfFields-15983]
	_ = x[ErrPathContainsEmptyElement-15998]
	_ = x[ErrOperatorWrongLenOfArgs-16020]
	_ = x[ErrFieldPathInvalidName-16410]
	_ = x[ErrGroupInvalidFieldPath-16872]
	_ = x[ErrGroupUndefinedVariable-17276]
	_ = x[ErrInvalidArg-28667]
	_ = x[ErrSliceFirstArg-28724]
	_ = x[ErrStageUnsetNoPath-31119]
	_ = x[ErrStageUnsetArrElementInvalidType-31120]
	_ = x[ErrStageUnsetInvalidType-31002]
	_ = x[ErrStageUnwindNoPath-28812]
	_ = x[ErrStageUnwindNoPrefix-28818]
	_ = x[ErrUnsetPathCollision-31249]
	_ = x[ErrUnsetPathOverwrite-31250]
	_ = x[ErrProjectionInEx-31253]
	_ = x[ErrProjectionExIn-31254]
	_ = x[ErrAggregatePositionalProject-31324]
	_ = x[ErrAggregateInvalidExpression-31325]
	_ = x[ErrWrongPositionalOperatorLocation-31394]
	_ = x[ErrExclusionPositionalProjection-31395]
	_ = x[ErrStageCountNonString-40156]
	_ = x[ErrStageCountNonEmptyString-40157]
	_ = x[ErrStageCountBadPrefix-40158]
	_ = x[ErrStageCountBadValue-40160]
	_ = x[ErrAddFieldsExpressionWrongAmountOfArgs-40181]
	_ = x[ErrStageGroupUnaryOperator-40237]
	_ = x[ErrStageGroupMultipleAccumulator-40238]
	_ = x[ErrStageGroupInvalidAccumulator-40234]
	_ = x[ErrStageInvalid-40323]
	_ = x[ErrEmptyFieldPath-40352]
	_ = x[ErrInvalidFieldPath-40353]
	_ = x[ErrMissingField-40414]
	_ = x[ErrFailedToParseInput-40415]
	_ = x[ErrCollStatsIsNotFirstStage-40602]
	_ = x[ErrFreeMonitoringDisabled-50840]
	_ = x[ErrUserAlreadyExists-51003]
	_ = x[ErrValueNegative-51024]
	_ = x[ErrRegexOptions-51075]
	_ = x[ErrRegexMissingParen-51091]
	_ = x[ErrBadRegexOption-51108]
	_ = x[ErrBadPositionalProjection-51246]
	_ = x[ErrElementMismatchPositionalProjection-51247]
	_ = x[ErrEmptySubProject-51270]
	_ = x[ErrEmptyProject-51272]
	_ = x[ErrDuplicateField-4822819]
	_ = x[ErrStageSkipBadValue-5107200]
	_ = x[ErrStageLimitInvalidArg-5107201]
	_ = x[ErrStageCollStatsInvalidArg-5447000]
	_ = x[ErrStageIndexedStringVectorDuplicate-7582300]
}

const _ErrorCode_name = "UnsetInternalErrorBadValueFailedToParseUserNotFoundUnauthorizedTypeMismatchAuthenticationFailedIllegalOperationNamespaceNotFoundIndexNotFoundPathNotViableConflictingUpdateOperatorsCursorNotFoundNamespaceExistsDollarPrefixedFieldNameInvalidIDEmptyFieldNameCommandNotFoundImmutableFieldCannotCreateIndexIndexAlreadyExistsInvalidOptionsInvalidNamespaceIndexOptionsConflictIndexKeySpecsConflictOperationFailedDocumentValidationFailureInvalidPipelineOperatorClientMetadataCannotBeMutatedInvalidIndexSpecificationOptionNotImplementedLocation10065Location11000Location15947Location15948Location15955Location15958Location15959Location15969Location15973Location15974Location15975Location15976Location15981Location15983Location15998Location16020Location16406Location16410Location16872Location17276Location28667Location28724Location28812Location28818Location31002Location31119Location31120Location31249Location31250Location31253Location31254Location31324Location31325Location31394Location31395Location40156Location40157Location40158Location40160Location40181Location40234Location40237Location40238Location40272Location40323Location40352Location40353Location40414Location40415Location40602Location50840Location51003Location51024Location51075Location51091Location51108Location51246Location51247Location51270Location51272Location4822819Location5107200Location5107201Location5447000Location7582300"

var _ErrorCode_map = map[ErrorCode]string{
	0:       _ErrorCode_name[0:5],
	1:       _ErrorCode_name[5:18],
	2:       _ErrorCode_name[18:26],
	9:       _ErrorCode_name[26:39],
	11:      _ErrorCode_name[39:51],
	13:      _ErrorCode_name[51:63],
	14:      _ErrorCode_name[63:75],
	18:      _ErrorCode_name[75:95],
	20:      _ErrorCode_name[95:111],
	26:      _ErrorCode_name[111:128],
	27:      _ErrorCode_name[128:141],
	28:      _ErrorCode_name[141:154],
	40:      _ErrorCode_name[154:180],
	43:      _ErrorCode_name[180:194],
	48:      _ErrorCode_name[194:209],
	52:      _ErrorCode_name[209:232],
	53:      _ErrorCode_name[232:241],
	56:      _ErrorCode_name[241:255],
	59:      _ErrorCode_name[255:270],
	66:      _ErrorCode_name[270:284],
	67:      _ErrorCode_name[284:301],
	68:      _ErrorCode_name[301:319],
	72:      _ErrorCode_name[319:333],
	73:      _ErrorCode_name[333:349],
	85:      _ErrorCode_name[349:369],
	86:      _ErrorCode_name[369:390],
	96:      _ErrorCode_name[390:405],
	121:     _ErrorCode_name[405:430],
	168:     _ErrorCode_name[430:453],
	186:     _ErrorCode_name[453:482],
	197:     _ErrorCode_name[482:513],
	238:     _ErrorCode_name[513:527],
	10065:   _ErrorCode_name[527:540],
	11000:   _ErrorCode_name[540:553],
	15947:   _ErrorCode_name[553:566],
	15948:   _ErrorCode_name[566:579],
	15955:   _ErrorCode_name[579:592],
	15958:   _ErrorCode_name[592:605],
	15959:   _ErrorCode_name[605:618],
	15969:   _ErrorCode_name[618:631],
	15973:   _ErrorCode_name[631:644],
	15974:   _ErrorCode_name[644:657],
	15975:   _ErrorCode_name[657:670],
	15976:   _ErrorCode_name[670:683],
	15981:   _ErrorCode_name[683:696],
	15983:   _ErrorCode_name[696:709],
	15998:   _ErrorCode_name[709:722],
	16020:   _ErrorCode_name[722:735],
	16406:   _ErrorCode_name[735:748],
	16410:   _ErrorCode_name[748:761],
	16872:   _ErrorCode_name[761:774],
	17276:   _ErrorCode_name[774:787],
	28667:   _ErrorCode_name[787:800],
	28724:   _ErrorCode_name[800:813],
	28812:   _ErrorCode_name[813:826],
	28818:   _ErrorCode_name[826:839],
	31002:   _ErrorCode_name[839:852],
	31119:   _ErrorCode_name[852:865],
	31120:   _ErrorCode_name[865:878],
	31249:   _ErrorCode_name[878:891],
	31250:   _ErrorCode_name[891:904],
	31253:   _ErrorCode_name[904:917],
	31254:   _ErrorCode_name[917:930],
	31324:   _ErrorCode_name[930:943],
	31325:   _ErrorCode_name[943:956],
	31394:   _ErrorCode_name[956:969],
	31395:   _ErrorCode_name[969:982],
	40156:   _ErrorCode_name[982:995],
	40157:   _ErrorCode_name[995:1008],
	40158:   _ErrorCode_name[1008:1021],
	40160:   _ErrorCode_name[1021:1034],
	40181:   _ErrorCode_name[1034:1047],
	40234:   _ErrorCode_name[1047:1060],
	40237:   _ErrorCode_name[1060:1073],
	40238:   _ErrorCode_name[1073:1086],
	40272:   _ErrorCode_name[1086:1099],
	40323:   _ErrorCode_name[1099:1112],
	40352:   _ErrorCode_name[1112:1125],
	40353:   _ErrorCode_name[1125:1138],
	40414:   _ErrorCode_name[1138:1151],
	40415:   _ErrorCode_name[1151:1164],
	40602:   _ErrorCode_name[1164:1177],
	50840:   _ErrorCode_name[1177:1190],
	51003:   _ErrorCode_name[1190:1203],
	51024:   _ErrorCode_name[1203:1216],
	51075:   _ErrorCode_name[1216:1229],
	51091:   _ErrorCode_name[1229:1242],
	51108:   _ErrorCode_name[1242:1255],
	51246:   _ErrorCode_name[1255:1268],
	51247:   _ErrorCode_name[1268:1281],
	51270:   _ErrorCode_name[1281:1294],
	51272:   _ErrorCode_name[1294:1307],
	4822819: _ErrorCode_name[1307:1322],
	5107200: _ErrorCode_name[1322:1337],
	5107201: _ErrorCode_name[1337:1352],
	5447000: _ErrorCode_name[1352:1367],
	7582300: _ErrorCode_name[1367:1382],
}

func (i ErrorCode) String() string {
	if str, ok := _ErrorCode_map[i]; ok {
		return str
	}
	return "ErrorCode(" + strconv.FormatInt(int64(i), 10) + ")"
}
