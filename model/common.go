package model

type (
	// LabelSet ラベルの集合
	LabelSet struct {
		Count int64 `xml:",attr"` // 要素数
		Label []Label
	}

	// Label ラベル
	Label struct {
		LabSid  int64  // 	ラベルSID
		LabName string // 	ラベル名
	}

	// GroupSet グループの集合
	GroupSet struct {
		Count int64   `xml:",attr"` // 要素数
		Group []Group // グループ情報
	}

	// Group グループ
	Group struct {
		GrpSid  int64  // グループSID
		GrpName string // グループ名
	}

	// PageSpecifyQuery ページだけを指定するクエリ（リクエストパラメータ）
	PageSpecifyQuery struct {
		Page int `query:"page"` // ページ番号
	}
)
