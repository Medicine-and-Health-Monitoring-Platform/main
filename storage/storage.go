package storage

import ( 
	"context"
	pb"main/genproto/health_analytics"
)
type IStorage interface {
	MedicalRecords() IMedicalRecordStorage
	LifestyleData() ILifestyleDataStorage
	WearableData() IWearableDataStorage
	Monitoring()  IMonitoringStorage
	Clouse()
}

type IMedicalRecordStorage interface{
	AddMedicalRecord(context.Context, *pb.AddMedicalRecordRequest)(*pb.AddMedicalRecordResponse, error)
	GetMedicalRecord(context.Context, *pb.GetMedicalRecordRequest)(*pb.GetMedicalRecordResponse, error)
	UpdateMedicalRecord(context.Context, *pb.UpdateMedicalRecordRequest)(*pb.UpdateMedicalRecordResponse, error)
	DeleteMedicalRecord(context.Context, *pb.DeleteMedicalRecordRequest)(*pb.DeleteMedicalRecordResponse, error)
	ListMedicalRecord(context.Context, *pb.ListMedicalRecordsRequest)(*pb.ListMedicalRecordsResponse, error)
}
 
type ILifestyleDataStorage interface{
	AddLifestyleData(context.Context, *pb.AddLifestyleDataRequest)(*pb.AddLifestyleDataResponse, error)
	GetLifestyleData(context.Context, *pb.GetLifestyleDataRequest)(*pb.GetLifestyleDataResponse, error)
	UpdateLifestyleData(context.Context, *pb.UpdateLifestyleDataRequest)(*pb.UpdateLifestyleDataResponse, error)
	DeleteLifestyleData(context.Context, *pb.DeleteLifestyleDataRequest)(*pb.DeleteLifestyleDataResponse, error)
}

type IWearableDataStorage interface{
	AddWearableData(context.Context, *pb.AddWearableDataRequest)(*pb.AddWearableDataResponse, error)
	GetWearableData(context.Context, *pb.GetWearableDataRequest)(*pb.GetWearableDataResponse, error)
	UpdateWearableData(context.Context, *pb.UpdateWearableDataRequest)(*pb.UpdateWearableDataResponse, error)
	DeleteWearableData(context.Context, *pb.DeleteWearableDataRequest)(*pb.DeleteWearableDataResponse, error)
}

type IMonitoringStorage interface{
	GenerateHealthRecommendations(context.Context, *pb.GenerateHealthRecommendationsRequest)(*pb.GenerateHealthRecommendationsResponse, error)
	GetRealtimeHealthMonitoring(context.Context, *pb.GetRealtimeHealthMonitoringRequest)(*pb.GetRealtimeHealthMonitoringResponse, error)
	GetDailyHealthSummary(context.Context, *pb.GetDailyHealthSummaryRequest)(*pb.GetDailyHealthSummaryResponse, error)
	GetWeeklyHealthSummary(context.Context, *pb.GetWeeklyHealthSummaryRequest)(*pb.GetWeeklyHealthSummaryResponse, error)
	CreateHealthMonitor(context.Context, *pb.CreateHealthMonitorReq)(error)
	GetHealthMonitor(context.Context, *pb.UserId) (*pb.GetHealthMonitorsRes, error)

}


