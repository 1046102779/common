package consts

const (
	// login or signin
	TYPE_LOGIN_PASSWORD              = 10
	TYPE_LOGIN_VERIFICATION_CODE     = 20
	TYPE__SIGN_IN__PASSWORD          = 10
	TYPE__SIGN_IN__VERIFICATION_CODE = 20

	// finance
	TYPE__FINANCE_ACCOUNT__CASH           = 10
	TYPE__FINANCE_ACCOUNT__BANK           = 20
	TYPE__FINANCE_ACCOUNT__WECHAT_PERSON  = 30
	TYPE__FINANCE_ACCOUNT__WECHAT_OFFICAL = 31
	TYPE__FINANCE_ACCOUNT__ALI_PAY        = 40

	// org
	TYPE__ORG__SHOP      = 10
	TYPE__ORG__WAREHOUSE = 20
	TYPE__ORG__COMPANY   = 30

	// user
	TYPE__USER__SHOP      = 10
	TYPE__USER__WAREHOUSE = 20
	TYPE__USER__COMPANY   = 30
	TYPE__USER__NORMAL    = 70

	// sex
	TYPE__SEX__MALE     = 1
	TYPE__SEX__FEMALE   = 2
	TYPE__SEX__UNKNOWED = 3

	// warehouse
	TYPE__WAREHOUSE__NORMAL = 10 //普通仓库

	// fabric pattern storage place
	TYPE__FABRIC_PATTERN_STORAGE__SHOP      = 10
	TYPE__FABRIC_PATTERN_STORAGE__WAREHOUSE = 20

	// official account permission
	TYPE_PERMISSION__OFFICIAL_ACCOUNT__MESSAGE_MANAGEMENT     = 1  // 消息管理权限
	TYPE_PERMISSION__OFFICIAL_ACCOUNT__USER_MANAGEMENT        = 2  // 用户管理权限
	TYPE_PERMISSION__OFFICIAL_ACCOUNT__ACCOUNT_SERVICE        = 3  // 帐号服务权限
	TYPE_PERMISSION__OFFICIAL_ACCOUNT__WEB_SERVICE            = 4  // 网页服务权限
	TYPE_PERMISSION__OFFICIAL_ACCOUNT__WECHAT_SHOP            = 5  // 微信小店权限
	TYPE_PERMISSION__OFFICIAL_ACCOUNT__MULTI_CUSTOMER_SERVICE = 6  // 微信多客服权限
	TYPE_PERMISSION__OFFICIAL_ACCOUNT__NOTIFICATION           = 7  // 群发与通知权限
	TYPE_PERMISSION__OFFICIAL_ACCOUNT__CARDS                  = 8  // 微信卡券权限
	TYPE_PERMISSION__OFFICIAL_ACCOUNT__SCANNING               = 9  // 微信扫一扫权限
	TYPE_PERMISSION__OFFICIAL_ACCOUNT__WECHAT_WIFI            = 10 // 微信连WIFI权限
	TYPE_PERMISSION__OFFICIAL_ACCOUNT__MATERIAL_MANAGEMENT    = 11 // 素材管理权限
	TYPE_PERMISSION__OFFICIAL_ACCOUNT__SHAKE_AROUND           = 12 // 微信摇周边权限
	TYPE_PERMISSION__OFFICIAL_ACCOUNT__WECHAT_STORE           = 13 // 微信门店权限
	TYPE_PERMISSION__OFFICIAL_ACCOUNT__WECHAT_PAY             = 14 // 微信支付权限
	TYPE_PERMISSION__OFFICIAL_ACCOUNT__CUSTOMIZE_MENU         = 15 // 自定义菜单权限

	// 微信用户类型
	TYPE_WECHAT__USER___BUSSINESS = 10 // 微信B端用户
	TYPE_WECHAT__USER__PERSONAL   = 20 // 微信C端用户

	// 手机号码状态
	TYPE_STATUS__MOBILE__UNCOMFIRMED = -10
	TYPE_STATUS__MOBILE__CONFIRMING  = 10
	TYPE_STATUS__MOBILE__CONFIRMED   = 20

	// 邮箱状态
	TYPE_STATUS__EMAIL__UNCOMFIRMED = -10
	TYPE_STATUS__EMAIL__CONFIRMING  = 10
	TYPE_STATUS__EMAIL__CONFIRMED   = 20

	// 邀请码状态
	TYPE__INVITATION_CODE__UNUSED = 10
	TYPE__INVITATION_CODE__USED   = 20

	// 面料存储的地点类型
	TYPE__INVENTORY__WAREHOUSE = 10 // 仓库
	TYPE__INVENTORY__SHOP      = 20 // 商铺

	// 微信支付方式
	TYPE_PAY__WECHAT__JSAPI  = "JSAPI"
	TYPE_PAY__WECHAT__NATIVE = "NATIVE"
	TYPE_PAY__WECHAT__APP    = "APP"

	// 支付业务场景类型
	TYPE_PAY_ENV__WECHAT__SMS_RECHARGE = "101" // 短信充值业务
	TYPE_PAY_ENV__WECHAT__FABRIC_ORDER = "102" // SaaS平台商家布匹交易

	// 微信是否已经绑定用户/客户账号
	TYPE__WECHAT_USER_BINDING__USER     = 10 // 绑定B端用户
	TYPE__WECHAT_USER_BINDING__CUSTOMER = 11 // 绑定C端用户
	TYPE__WECHAT_USER_BINDING__NO       = 20

	// 库存启用状态: 10: 库存已启用；20：库存未启用
	TYPE__STORAGE__ENABLED = 10
	TYPE__STORAGE__UNABLED = 20

	// 微信支付开启状态
	TYPE_PAY_STATUS__WECHAT__STARTED = 10 // 已开启
	TYPE_PAY_STATUS__WECHAT__UNSTART = 20 // 未开启

	// 模板消息类型
	TYPE__MESSAGE_TEMPLATE__WAITING_TO_OUTS__WAREHOUSE_KEEPER = 10 // 仓管员-接收待出库提醒消息
	TYPE__MESSAGE_TEMPLATE__OUTS__SALER_OR_SHOP_MANAGER       = 20 // 销售员或店长-接收已出库提醒
	TYPE__MESSAGE_TEMPLATE__LOGISTICS__SALER_OR_SHOP_MANAGER  = 30 // 销售员或店长-接收已物流提醒
	TYPE__MESSAGE_TEMPLATE__PLACE_ORDER__CUSTOMER             = 40 // 客户-下单成功时提醒
	TYPE__MESSAGE_TEMPLATE__PRICE__CUSTOMER                   = 50 // 客户-总价格确定时接收提醒
	TYPE__MESSAGE_TEMPLATE__ACCEPTION__CUSTOMER               = 60 // 客户-待客户验收状态时接收提醒
	TYPE__MESSAGE_TEMPLATE__STATIS__GENERAL_MANAGER           = 70 // 总经理-每晚8点接收当天营业额统计情况的提醒

	// Customer mobile  sale order's order status
	TYPE__CUSTOMER_MOBILE__SALE_ORDER__ORDER_STATUS__TO_BE_RECEIVING = 10 // C端客户待收货
	TYPE__CUSTOMER_MOBILE__SALE_ORDER__ORDER_STATUS__RECEIVED        = 20 // C端客户已收货

	// customer mobile sale order pay status
	TYPE_CUSTOMER_MOBILE__SALE_ORDER__PAY_STATUS__TO_BE_PAY                = 10 // 待支付
	TYPE_CUSTOMER_MOBILE__SALE_ORDER__PAY_STATUS__TO_BE_PAY_BALANCE_AMOUNT = 10 // 待付尾款
	TYPE_CUSTOMER_MOBILE__SALE_ORDER__PAY_STATUS__PAYED                    = 10 // 已支付
)
