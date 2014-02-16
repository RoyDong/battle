package model

import (
    pt "github.com/roydong/potato"
    "github.com/roydong/potato/orm"
    "github.com/roydong/potato/lib"
)

var (
    Conf *lib.Tree

    MapModel *mapModel

    UserModel *userModel

    RoleModel *roleModel

    UserRoleModel *userRoleModel

    BaseModel *baseModel

    CenterModel *centerModel

    LabModel *labModel

    ArmoryModel *armoryModel

    SupplyModel *supplyModel

    StopeModel *stopeModel

    StorageModel *storageModel
)

func init() {
    pt.AddHandler("after_init", func(args ...interface{}) {
        Conf = lib.NewTree()
        if e := Conf.LoadYaml(pt.ConfDir + "model.yml", true); e != nil {
            pt.Logger.Fatal(e)
        }

        if v, has := Conf.Int("map.resource_max"); has {
            ResourceMax = v
        }
        if v, has := Conf.Int("map.resource_min"); has {
            ResourceMin = v
        }
        if v, has := Conf.Int("map.resource_rate"); has {
            ResourceRate = v
        }
        if v, has := Conf.Int("map.energe_rate"); has {
            EnergyRate = v
        }
    })

    pt.AddHandler("after_init", func(args ...interface{}) {
        MapModel = newMapModel()

        UserModel = &userModel{orm.NewModel("user", &User{})}

        RoleModel = &roleModel{orm.NewModel("role", &Role{})}

        UserRoleModel = &userRoleModel{orm.NewModel("user_role", &UserRole{})}
        BaseModel = &baseModel{orm.NewModel("base", &Base{})}

        CenterModel = &centerModel{orm.NewModel("main", &Center{})}

        LabModel = &labModel{orm.NewModel("lab", &Lab{})}

        ArmoryModel = &armoryModel{orm.NewModel("armory", &Armory{})}

        SupplyModel = &supplyModel{orm.NewModel("supply", &Supply{})}

        StopeModel = &stopeModel{orm.NewModel("stope", &Stope{})}

        StorageModel = &storageModel{orm.NewModel("storage", &Storage{})}
    })
}
