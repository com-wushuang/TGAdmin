package request

import "gin-vue-admin/model"

type JobSearch struct{
    model.Job
    PageInfo
}