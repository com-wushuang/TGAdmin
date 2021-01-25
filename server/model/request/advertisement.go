package request

import "gin-vue-admin/model"

type AdvertisementSearch struct{
    model.Advertisement
    PageInfo
}