// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: product_service/product.proto

package product_service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ProductServiceClient is the client API for ProductService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductServiceClient interface {
	AddProduct(ctx context.Context, in *AddProductRequest, opts ...grpc.CallOption) (*AddProductResponse, error)
	EditProduct(ctx context.Context, in *EditProductRequest, opts ...grpc.CallOption) (*EditProductResponse, error)
	GetProduct(ctx context.Context, in *GetProductRequest, opts ...grpc.CallOption) (*GetProductResponse, error)
	GetProducts(ctx context.Context, in *GetProductsRequest, opts ...grpc.CallOption) (*GetProductsResponse, error)
	DeleteProduct(ctx context.Context, in *DeleteProductRequest, opts ...grpc.CallOption) (*DeleteProductResponse, error)
	SearchProducts(ctx context.Context, in *SearchProductsRequest, opts ...grpc.CallOption) (*SearchProductsResponse, error)
	AddRating(ctx context.Context, in *AddRatingRequest, opts ...grpc.CallOption) (*AddRatingResponse, error)
	GetRatings(ctx context.Context, in *GetRatingsRequest, opts ...grpc.CallOption) (*GetRatingsResponse, error)
	PlaceOrder(ctx context.Context, in *PlaceOrderRequest, opts ...grpc.CallOption) (*PlaceOrderResponse, error)
	CancelOrder(ctx context.Context, in *CancelOrderRequest, opts ...grpc.CallOption) (*CancelOrderResponse, error)
	UpdateOrderStatus(ctx context.Context, in *UpdateOrderStatusRequest, opts ...grpc.CallOption) (*UpdateOrderStatusResponse, error)
	GetOrders(ctx context.Context, in *GetOrdersRequest, opts ...grpc.CallOption) (*GetOrdersResponse, error)
	GetOrder(ctx context.Context, in *GetOrderRequest, opts ...grpc.CallOption) (*GetOrderResponse, error)
	PayOrder(ctx context.Context, in *PayOrderRequest, opts ...grpc.CallOption) (*PayOrderResponse, error)
	CheckPaymentStatus(ctx context.Context, in *CheckPaymentStatusRequest, opts ...grpc.CallOption) (*CheckPaymentStatusResponse, error)
	UpdateShippingInfo(ctx context.Context, in *UpdateShippingInfoRequest, opts ...grpc.CallOption) (*UpdateShippingInfoResponse, error)
	AddArtisanCategory(ctx context.Context, in *AddArtisanCategoryRequest, opts ...grpc.CallOption) (*AddArtisanCategoryResponse, error)
	AddProductCategory(ctx context.Context, in *AddProductCategoryRequest, opts ...grpc.CallOption) (*AddProductCategoryResponse, error)
	GetStatistics(ctx context.Context, in *GetStatisticsRequest, opts ...grpc.CallOption) (*GetStatisticsResponse, error)
	GetUserActivity(ctx context.Context, in *GetUserActivityRequest, opts ...grpc.CallOption) (*GetUserActivityResponse, error)
	GetRecommendations(ctx context.Context, in *GetRecommendationsRequest, opts ...grpc.CallOption) (*GetRecommendationsResponse, error)
	GetArtisanRankings(ctx context.Context, in *GetArtisanRankingsRequest, opts ...grpc.CallOption) (*GetArtisanRankingsResponse, error)
}

type productServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProductServiceClient(cc grpc.ClientConnInterface) ProductServiceClient {
	return &productServiceClient{cc}
}

func (c *productServiceClient) AddProduct(ctx context.Context, in *AddProductRequest, opts ...grpc.CallOption) (*AddProductResponse, error) {
	out := new(AddProductResponse)
	err := c.cc.Invoke(ctx, "/product_service.ProductService/AddProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) EditProduct(ctx context.Context, in *EditProductRequest, opts ...grpc.CallOption) (*EditProductResponse, error) {
	out := new(EditProductResponse)
	err := c.cc.Invoke(ctx, "/product_service.ProductService/EditProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) GetProduct(ctx context.Context, in *GetProductRequest, opts ...grpc.CallOption) (*GetProductResponse, error) {
	out := new(GetProductResponse)
	err := c.cc.Invoke(ctx, "/product_service.ProductService/GetProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) GetProducts(ctx context.Context, in *GetProductsRequest, opts ...grpc.CallOption) (*GetProductsResponse, error) {
	out := new(GetProductsResponse)
	err := c.cc.Invoke(ctx, "/product_service.ProductService/GetProducts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) DeleteProduct(ctx context.Context, in *DeleteProductRequest, opts ...grpc.CallOption) (*DeleteProductResponse, error) {
	out := new(DeleteProductResponse)
	err := c.cc.Invoke(ctx, "/product_service.ProductService/DeleteProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) SearchProducts(ctx context.Context, in *SearchProductsRequest, opts ...grpc.CallOption) (*SearchProductsResponse, error) {
	out := new(SearchProductsResponse)
	err := c.cc.Invoke(ctx, "/product_service.ProductService/SearchProducts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) AddRating(ctx context.Context, in *AddRatingRequest, opts ...grpc.CallOption) (*AddRatingResponse, error) {
	out := new(AddRatingResponse)
	err := c.cc.Invoke(ctx, "/product_service.ProductService/AddRating", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) GetRatings(ctx context.Context, in *GetRatingsRequest, opts ...grpc.CallOption) (*GetRatingsResponse, error) {
	out := new(GetRatingsResponse)
	err := c.cc.Invoke(ctx, "/product_service.ProductService/GetRatings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) PlaceOrder(ctx context.Context, in *PlaceOrderRequest, opts ...grpc.CallOption) (*PlaceOrderResponse, error) {
	out := new(PlaceOrderResponse)
	err := c.cc.Invoke(ctx, "/product_service.ProductService/PlaceOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) CancelOrder(ctx context.Context, in *CancelOrderRequest, opts ...grpc.CallOption) (*CancelOrderResponse, error) {
	out := new(CancelOrderResponse)
	err := c.cc.Invoke(ctx, "/product_service.ProductService/CancelOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) UpdateOrderStatus(ctx context.Context, in *UpdateOrderStatusRequest, opts ...grpc.CallOption) (*UpdateOrderStatusResponse, error) {
	out := new(UpdateOrderStatusResponse)
	err := c.cc.Invoke(ctx, "/product_service.ProductService/UpdateOrderStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) GetOrders(ctx context.Context, in *GetOrdersRequest, opts ...grpc.CallOption) (*GetOrdersResponse, error) {
	out := new(GetOrdersResponse)
	err := c.cc.Invoke(ctx, "/product_service.ProductService/GetOrders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) GetOrder(ctx context.Context, in *GetOrderRequest, opts ...grpc.CallOption) (*GetOrderResponse, error) {
	out := new(GetOrderResponse)
	err := c.cc.Invoke(ctx, "/product_service.ProductService/GetOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) PayOrder(ctx context.Context, in *PayOrderRequest, opts ...grpc.CallOption) (*PayOrderResponse, error) {
	out := new(PayOrderResponse)
	err := c.cc.Invoke(ctx, "/product_service.ProductService/PayOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) CheckPaymentStatus(ctx context.Context, in *CheckPaymentStatusRequest, opts ...grpc.CallOption) (*CheckPaymentStatusResponse, error) {
	out := new(CheckPaymentStatusResponse)
	err := c.cc.Invoke(ctx, "/product_service.ProductService/CheckPaymentStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) UpdateShippingInfo(ctx context.Context, in *UpdateShippingInfoRequest, opts ...grpc.CallOption) (*UpdateShippingInfoResponse, error) {
	out := new(UpdateShippingInfoResponse)
	err := c.cc.Invoke(ctx, "/product_service.ProductService/UpdateShippingInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) AddArtisanCategory(ctx context.Context, in *AddArtisanCategoryRequest, opts ...grpc.CallOption) (*AddArtisanCategoryResponse, error) {
	out := new(AddArtisanCategoryResponse)
	err := c.cc.Invoke(ctx, "/product_service.ProductService/AddArtisanCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) AddProductCategory(ctx context.Context, in *AddProductCategoryRequest, opts ...grpc.CallOption) (*AddProductCategoryResponse, error) {
	out := new(AddProductCategoryResponse)
	err := c.cc.Invoke(ctx, "/product_service.ProductService/AddProductCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) GetStatistics(ctx context.Context, in *GetStatisticsRequest, opts ...grpc.CallOption) (*GetStatisticsResponse, error) {
	out := new(GetStatisticsResponse)
	err := c.cc.Invoke(ctx, "/product_service.ProductService/GetStatistics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) GetUserActivity(ctx context.Context, in *GetUserActivityRequest, opts ...grpc.CallOption) (*GetUserActivityResponse, error) {
	out := new(GetUserActivityResponse)
	err := c.cc.Invoke(ctx, "/product_service.ProductService/GetUserActivity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) GetRecommendations(ctx context.Context, in *GetRecommendationsRequest, opts ...grpc.CallOption) (*GetRecommendationsResponse, error) {
	out := new(GetRecommendationsResponse)
	err := c.cc.Invoke(ctx, "/product_service.ProductService/GetRecommendations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) GetArtisanRankings(ctx context.Context, in *GetArtisanRankingsRequest, opts ...grpc.CallOption) (*GetArtisanRankingsResponse, error) {
	out := new(GetArtisanRankingsResponse)
	err := c.cc.Invoke(ctx, "/product_service.ProductService/GetArtisanRankings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductServiceServer is the server API for ProductService service.
// All implementations must embed UnimplementedProductServiceServer
// for forward compatibility
type ProductServiceServer interface {
	AddProduct(context.Context, *AddProductRequest) (*AddProductResponse, error)
	EditProduct(context.Context, *EditProductRequest) (*EditProductResponse, error)
	GetProduct(context.Context, *GetProductRequest) (*GetProductResponse, error)
	GetProducts(context.Context, *GetProductsRequest) (*GetProductsResponse, error)
	DeleteProduct(context.Context, *DeleteProductRequest) (*DeleteProductResponse, error)
	SearchProducts(context.Context, *SearchProductsRequest) (*SearchProductsResponse, error)
	AddRating(context.Context, *AddRatingRequest) (*AddRatingResponse, error)
	GetRatings(context.Context, *GetRatingsRequest) (*GetRatingsResponse, error)
	PlaceOrder(context.Context, *PlaceOrderRequest) (*PlaceOrderResponse, error)
	CancelOrder(context.Context, *CancelOrderRequest) (*CancelOrderResponse, error)
	UpdateOrderStatus(context.Context, *UpdateOrderStatusRequest) (*UpdateOrderStatusResponse, error)
	GetOrders(context.Context, *GetOrdersRequest) (*GetOrdersResponse, error)
	GetOrder(context.Context, *GetOrderRequest) (*GetOrderResponse, error)
	PayOrder(context.Context, *PayOrderRequest) (*PayOrderResponse, error)
	CheckPaymentStatus(context.Context, *CheckPaymentStatusRequest) (*CheckPaymentStatusResponse, error)
	UpdateShippingInfo(context.Context, *UpdateShippingInfoRequest) (*UpdateShippingInfoResponse, error)
	AddArtisanCategory(context.Context, *AddArtisanCategoryRequest) (*AddArtisanCategoryResponse, error)
	AddProductCategory(context.Context, *AddProductCategoryRequest) (*AddProductCategoryResponse, error)
	GetStatistics(context.Context, *GetStatisticsRequest) (*GetStatisticsResponse, error)
	GetUserActivity(context.Context, *GetUserActivityRequest) (*GetUserActivityResponse, error)
	GetRecommendations(context.Context, *GetRecommendationsRequest) (*GetRecommendationsResponse, error)
	GetArtisanRankings(context.Context, *GetArtisanRankingsRequest) (*GetArtisanRankingsResponse, error)
	mustEmbedUnimplementedProductServiceServer()
}

// UnimplementedProductServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProductServiceServer struct {
}

func (UnimplementedProductServiceServer) AddProduct(context.Context, *AddProductRequest) (*AddProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddProduct not implemented")
}
func (UnimplementedProductServiceServer) EditProduct(context.Context, *EditProductRequest) (*EditProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditProduct not implemented")
}
func (UnimplementedProductServiceServer) GetProduct(context.Context, *GetProductRequest) (*GetProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProduct not implemented")
}
func (UnimplementedProductServiceServer) GetProducts(context.Context, *GetProductsRequest) (*GetProductsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProducts not implemented")
}
func (UnimplementedProductServiceServer) DeleteProduct(context.Context, *DeleteProductRequest) (*DeleteProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProduct not implemented")
}
func (UnimplementedProductServiceServer) SearchProducts(context.Context, *SearchProductsRequest) (*SearchProductsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchProducts not implemented")
}
func (UnimplementedProductServiceServer) AddRating(context.Context, *AddRatingRequest) (*AddRatingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddRating not implemented")
}
func (UnimplementedProductServiceServer) GetRatings(context.Context, *GetRatingsRequest) (*GetRatingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRatings not implemented")
}
func (UnimplementedProductServiceServer) PlaceOrder(context.Context, *PlaceOrderRequest) (*PlaceOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlaceOrder not implemented")
}
func (UnimplementedProductServiceServer) CancelOrder(context.Context, *CancelOrderRequest) (*CancelOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelOrder not implemented")
}
func (UnimplementedProductServiceServer) UpdateOrderStatus(context.Context, *UpdateOrderStatusRequest) (*UpdateOrderStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOrderStatus not implemented")
}
func (UnimplementedProductServiceServer) GetOrders(context.Context, *GetOrdersRequest) (*GetOrdersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrders not implemented")
}
func (UnimplementedProductServiceServer) GetOrder(context.Context, *GetOrderRequest) (*GetOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrder not implemented")
}
func (UnimplementedProductServiceServer) PayOrder(context.Context, *PayOrderRequest) (*PayOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PayOrder not implemented")
}
func (UnimplementedProductServiceServer) CheckPaymentStatus(context.Context, *CheckPaymentStatusRequest) (*CheckPaymentStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckPaymentStatus not implemented")
}
func (UnimplementedProductServiceServer) UpdateShippingInfo(context.Context, *UpdateShippingInfoRequest) (*UpdateShippingInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateShippingInfo not implemented")
}
func (UnimplementedProductServiceServer) AddArtisanCategory(context.Context, *AddArtisanCategoryRequest) (*AddArtisanCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddArtisanCategory not implemented")
}
func (UnimplementedProductServiceServer) AddProductCategory(context.Context, *AddProductCategoryRequest) (*AddProductCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddProductCategory not implemented")
}
func (UnimplementedProductServiceServer) GetStatistics(context.Context, *GetStatisticsRequest) (*GetStatisticsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatistics not implemented")
}
func (UnimplementedProductServiceServer) GetUserActivity(context.Context, *GetUserActivityRequest) (*GetUserActivityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserActivity not implemented")
}
func (UnimplementedProductServiceServer) GetRecommendations(context.Context, *GetRecommendationsRequest) (*GetRecommendationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRecommendations not implemented")
}
func (UnimplementedProductServiceServer) GetArtisanRankings(context.Context, *GetArtisanRankingsRequest) (*GetArtisanRankingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArtisanRankings not implemented")
}
func (UnimplementedProductServiceServer) mustEmbedUnimplementedProductServiceServer() {}

// UnsafeProductServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductServiceServer will
// result in compilation errors.
type UnsafeProductServiceServer interface {
	mustEmbedUnimplementedProductServiceServer()
}

func RegisterProductServiceServer(s grpc.ServiceRegistrar, srv ProductServiceServer) {
	s.RegisterService(&ProductService_ServiceDesc, srv)
}

func _ProductService_AddProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).AddProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product_service.ProductService/AddProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).AddProduct(ctx, req.(*AddProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_EditProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).EditProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product_service.ProductService/EditProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).EditProduct(ctx, req.(*EditProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_GetProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).GetProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product_service.ProductService/GetProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).GetProduct(ctx, req.(*GetProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_GetProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).GetProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product_service.ProductService/GetProducts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).GetProducts(ctx, req.(*GetProductsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_DeleteProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).DeleteProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product_service.ProductService/DeleteProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).DeleteProduct(ctx, req.(*DeleteProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_SearchProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchProductsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).SearchProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product_service.ProductService/SearchProducts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).SearchProducts(ctx, req.(*SearchProductsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_AddRating_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRatingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).AddRating(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product_service.ProductService/AddRating",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).AddRating(ctx, req.(*AddRatingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_GetRatings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRatingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).GetRatings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product_service.ProductService/GetRatings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).GetRatings(ctx, req.(*GetRatingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_PlaceOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlaceOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).PlaceOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product_service.ProductService/PlaceOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).PlaceOrder(ctx, req.(*PlaceOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_CancelOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).CancelOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product_service.ProductService/CancelOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).CancelOrder(ctx, req.(*CancelOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_UpdateOrderStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOrderStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).UpdateOrderStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product_service.ProductService/UpdateOrderStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).UpdateOrderStatus(ctx, req.(*UpdateOrderStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_GetOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).GetOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product_service.ProductService/GetOrders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).GetOrders(ctx, req.(*GetOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_GetOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).GetOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product_service.ProductService/GetOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).GetOrder(ctx, req.(*GetOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_PayOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PayOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).PayOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product_service.ProductService/PayOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).PayOrder(ctx, req.(*PayOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_CheckPaymentStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckPaymentStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).CheckPaymentStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product_service.ProductService/CheckPaymentStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).CheckPaymentStatus(ctx, req.(*CheckPaymentStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_UpdateShippingInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateShippingInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).UpdateShippingInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product_service.ProductService/UpdateShippingInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).UpdateShippingInfo(ctx, req.(*UpdateShippingInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_AddArtisanCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddArtisanCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).AddArtisanCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product_service.ProductService/AddArtisanCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).AddArtisanCategory(ctx, req.(*AddArtisanCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_AddProductCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddProductCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).AddProductCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product_service.ProductService/AddProductCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).AddProductCategory(ctx, req.(*AddProductCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_GetStatistics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStatisticsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).GetStatistics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product_service.ProductService/GetStatistics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).GetStatistics(ctx, req.(*GetStatisticsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_GetUserActivity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserActivityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).GetUserActivity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product_service.ProductService/GetUserActivity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).GetUserActivity(ctx, req.(*GetUserActivityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_GetRecommendations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRecommendationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).GetRecommendations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product_service.ProductService/GetRecommendations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).GetRecommendations(ctx, req.(*GetRecommendationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_GetArtisanRankings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetArtisanRankingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).GetArtisanRankings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product_service.ProductService/GetArtisanRankings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).GetArtisanRankings(ctx, req.(*GetArtisanRankingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProductService_ServiceDesc is the grpc.ServiceDesc for ProductService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProductService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "product_service.ProductService",
	HandlerType: (*ProductServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddProduct",
			Handler:    _ProductService_AddProduct_Handler,
		},
		{
			MethodName: "EditProduct",
			Handler:    _ProductService_EditProduct_Handler,
		},
		{
			MethodName: "GetProduct",
			Handler:    _ProductService_GetProduct_Handler,
		},
		{
			MethodName: "GetProducts",
			Handler:    _ProductService_GetProducts_Handler,
		},
		{
			MethodName: "DeleteProduct",
			Handler:    _ProductService_DeleteProduct_Handler,
		},
		{
			MethodName: "SearchProducts",
			Handler:    _ProductService_SearchProducts_Handler,
		},
		{
			MethodName: "AddRating",
			Handler:    _ProductService_AddRating_Handler,
		},
		{
			MethodName: "GetRatings",
			Handler:    _ProductService_GetRatings_Handler,
		},
		{
			MethodName: "PlaceOrder",
			Handler:    _ProductService_PlaceOrder_Handler,
		},
		{
			MethodName: "CancelOrder",
			Handler:    _ProductService_CancelOrder_Handler,
		},
		{
			MethodName: "UpdateOrderStatus",
			Handler:    _ProductService_UpdateOrderStatus_Handler,
		},
		{
			MethodName: "GetOrders",
			Handler:    _ProductService_GetOrders_Handler,
		},
		{
			MethodName: "GetOrder",
			Handler:    _ProductService_GetOrder_Handler,
		},
		{
			MethodName: "PayOrder",
			Handler:    _ProductService_PayOrder_Handler,
		},
		{
			MethodName: "CheckPaymentStatus",
			Handler:    _ProductService_CheckPaymentStatus_Handler,
		},
		{
			MethodName: "UpdateShippingInfo",
			Handler:    _ProductService_UpdateShippingInfo_Handler,
		},
		{
			MethodName: "AddArtisanCategory",
			Handler:    _ProductService_AddArtisanCategory_Handler,
		},
		{
			MethodName: "AddProductCategory",
			Handler:    _ProductService_AddProductCategory_Handler,
		},
		{
			MethodName: "GetStatistics",
			Handler:    _ProductService_GetStatistics_Handler,
		},
		{
			MethodName: "GetUserActivity",
			Handler:    _ProductService_GetUserActivity_Handler,
		},
		{
			MethodName: "GetRecommendations",
			Handler:    _ProductService_GetRecommendations_Handler,
		},
		{
			MethodName: "GetArtisanRankings",
			Handler:    _ProductService_GetArtisanRankings_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "product_service/product.proto",
}
