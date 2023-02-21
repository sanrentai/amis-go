package amis

// Grid
func Grid2d(opts ...opt) map[string]interface{} {
	return newCompent("grid-2d", opts...)
}

// 外层 Dom 的类名
func Grid2d_gridClassName(p string) opt {
	return func(o map[string]interface{}) {
		o["gridClassName"] = p
	}
}

// 格子间距，包括水平和垂直
func Grid2d_gap(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["gap"] = p
	}
}

// 格子水平划分为几个区域
func Grid2d_cols(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["cols"] = p
	}
}

// 每个格子默认垂直高度
func Grid2d_rowHeight(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["rowHeight"] = p
	}
}

// 格子垂直间距
func Grid2d_rowGap(p interface{}) opt {
	return func(o map[string]interface{}) {
		o["rowGap"] = p
	}
}

// 格子集合
func Grid2d_grids(p ...interface{}) opt {
	return func(o map[string]interface{}) {
		o["grids"] = p
	}
}

// // 格子可以是其他渲染器
// func Grid2d_grids[x](p interface{}) opt {
// 	return func(o map[string]interface{}) {
// 		o["grids[x]"] = p
// 	}
// }

// // 格子起始位置的横坐标
// func Grid2d_grids[x].x(p interface{}) opt {
// 	return func(o map[string]interface{}) {
// 		o["grids[x].x"] = p
// 	}
// }

// // 格子起始位置的纵坐标
// func Grid2d_grids[x].y(p interface{}) opt {
// 	return func(o map[string]interface{}) {
// 		o["grids[x].y"] = p
// 	}
// }

// // 格子横跨几个宽度
// func Grid2d_grids[x].w(p interface{}) opt {
// 	return func(o map[string]interface{}) {
// 		o["grids[x].w"] = p
// 	}
// }

// // 格子横跨几个高度
// func Grid2d_grids[x].h(p interface{}) opt {
// 	return func(o map[string]interface{}) {
// 		o["grids[x].h"] = p
// 	}
// }

// // 格子所在列的宽度
// func Grid2d_grids[x].width(p interface{}) opt {
// 	return func(o map[string]interface{}) {
// 		o["grids[x].width"] = p
// 	}
// }

// // 格子所在行的高度，可以设置 auto
// func Grid2d_grids[x].height(p interface{}) opt {
// 	return func(o map[string]interface{}) {
// 		o["grids[x].height"] = p
// 	}
// }

// // 格子内容水平布局
// func Grid2d_grids[x].align(p interface{}) opt {
// 	return func(o map[string]interface{}) {
// 		o["grids[x].align"] = p
// 	}
// }

// // 格子内容垂直布局
// func Grid2d_grids[x].valign(p interface{}) opt {
// 	return func(o map[string]interface{}) {
// 		o["grids[x].valign"] = p
// 	}
// }
