package request

import "gin-vue-admin/model"

type AccountSearch struct{
    model.Account
    PageInfo
}