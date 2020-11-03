export default {
    state: {
        interface_wx_visible: false,
        task_wx_visible: false,
        interface_qqmusic_visible: false,
        task_qqmusic_visible: false,
        interface_tixing_visible: false,
        task_tixing_visible: false,
        interface_xiangce_visible: false,
        task_xiangce_visible: false,
        interface_rili_visible: false,
        task_rili_visible: false,
        interface_tianqi_visible: false,
        task_tianqi_visible: false,
        interface_clock_visible: false,
        task_clock_visible: false,
        interface_jisuanqi_visible: false,
        task_jisuanqi_visible: false,
        interface_shipin_visible: false,
        task_shipin_visible: false,
        interface_wenjianjia_visible: false,
        task_wenjianjia_visible: false,
        interface_shop_visible: false,
        task_shop_visible: false,
        bgc: require('../assets/desktop/bg.jpg'),
        remindList: [
            {
                content: '下午两点开会',
                date: '2019-12-01',
                done: true
            },
            {
                content: '明天吃烤鸭',
                date: '2019-12-02',
                done: true
            },
            {
                content: '后天吃大盘鸡',
                date: '2019-12-02',
                done: false
            },
            {
                content: '周末去爬山',
                date: '2019-12-04',
                done: false
            }
        ]

    },
    mutations: {
        updateBgc(state, newbgc) {
            state.bgc = newbgc
        },
        updateRemindList(state, newList) {
            state.remindList = newList
        },
        changeInterfaceWxVisible(state) {
            state.interface_wx_visible = !state.interface_wx_visible
        },
        changeTaskWxVisible(state) {
            state.task_wx_visible = !state.task_wx_visible
        },
        changeInterfaceQQmusicVisible(state) {
            state.interface_qqmusic_visible = !state.interface_qqmusic_visible
        },
        changeTaskQQmusicVisible(state) {
            state.task_qqmusic_visible = !state.task_qqmusic_visible
        },
        changeInterfaceTixingVisible(state) {
            state.interface_tixing_visible = !state.interface_tixing_visible
        },
        changeTaskTixingVisible(state) {
            state.task_tixing_visible = !state.task_tixing_visible
        },
        changeInterfaceXiangceVisible(state) {
            state.interface_xiangce_visible = !state.interface_xiangce_visible
        },
        changeTaskXiangceVisible(state) {
            state.task_xiangce_visible = !state.task_xiangce_visible
        },
        changeInterfaceRiliVisible(state) {
            state.interface_rili_visible = !state.interface_rili_visible
        },
        changeTaskRiliVisible(state) {
            state.task_rili_visible = !state.task_rili_visible
        },
        changeInterfaceTianqiVisible(state) {
            state.interface_tianqi_visible = !state.interface_tianqi_visible
        },
        changeTaskTianqiVisible(state) {
            state.task_tianqi_visible = !state.task_tianqi_visible
        },
        changeInterfaceClockVisible(state) {
            state.interface_clock_visible = !state.interface_clock_visible
        },
        changeTaskClockVisible(state) {
            state.task_clock_visible = !state.task_clock_visible
        },
        changeInterfaceJisuanqiVisible(state) {
            state.interface_jisuanqi_visible = !state.interface_jisuanqi_visible
        },
        changeTaskJisuanqiVisible(state) {
            state.task_jisuanqi_visible = !state.task_jisuanqi_visible
        },
        changeInterfaceShipinVisible(state) {
            state.interface_shipin_visible = !state.interface_shipin_visible
        },
        changeTaskShipinVisible(state) {
            state.task_shipin_visible = !state.task_shipin_visible
        },
        changeInterfaceWenjianjiaVisible(state) {
            state.interface_wenjianjia_visible = !state.interface_wenjianjia_visible
        },
        changeTaskWenjianjiaVisible(state) {
            state.task_wenjianjia_visible = !state.task_wenjianjia_visible
        },
        changeInterfaceShopVisible(state) {
            state.interface_shop_visible = !state.interface_shop_visible
        },
        changeTaskShopVisible(state) {
            state.task_shop_visible = !state.task_shop_visible
        }
    },
    actions: {
        updateRemindList(context, newlist) {
            context.commit('updateRemindList', newlist)
        },
        updateBgc(context, newBgc) {
            context.commit('updateBgc', newBgc)
        }
    },
    modules: {}
}
