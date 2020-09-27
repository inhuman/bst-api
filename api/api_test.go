package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/inhuman/bst-api/mocks"
	"github.com/stretchr/testify/mock"
	"net/http"
	"testing"
)

func TestContainer_GetByKeyOk(t *testing.T) {

	rootNodeMock := &mocks.TreeNode{}

	bstContainerMock := &mocks.Container{}
	bstContainerMock.On("GetRoot").Return(rootNodeMock)
	bstContainerMock.On("Find", rootNodeMock, 77).Return("77-value")

	ginMock := &mocks.GinContext{}
	ginMock.On("Query", "val").Return("77")
	ginMock.On("AbortWithStatusJSON", http.StatusOK, gin.H{
		"value": "77-value",
	})

	apiContainer := &Container{
		BstContainer: bstContainerMock,
	}

	apiContainer.GetByKey(ginMock)

	bstContainerMock.AssertExpectations(t)
	ginMock.AssertExpectations(t)
}

func TestContainer_GetByKeyValueNotFound(t *testing.T) {

	rootNodeMock := &mocks.TreeNode{}

	bstContainerMock := &mocks.Container{}
	bstContainerMock.On("GetRoot").Return(rootNodeMock)
	bstContainerMock.On("Find", rootNodeMock, 77).Return(nil)

	ginMock := &mocks.GinContext{}
	ginMock.On("Query", "val").Return("77")
	ginMock.On("AbortWithStatusJSON", http.StatusNotFound, gin.H{
		"message": "value for key 77 not found",
	})

	apiContainer := &Container{
		BstContainer: bstContainerMock,
	}

	apiContainer.GetByKey(ginMock)

	bstContainerMock.AssertExpectations(t)
	ginMock.AssertExpectations(t)
}

func TestContainer_GetByKeyIncorrectValue(t *testing.T) {

	bstContainerMock := &mocks.Container{}

	ginMock := &mocks.GinContext{}
	ginMock.On("Query", "val").Return("incorrect value")
	ginMock.On("AbortWithStatusJSON", http.StatusInternalServerError, mock.Anything)

	apiContainer := &Container{
		BstContainer: bstContainerMock,
	}

	apiContainer.GetByKey(ginMock)

	bstContainerMock.AssertExpectations(t)
	ginMock.AssertExpectations(t)
}

func TestContainer_DeleteByKeyOk(t *testing.T) {

	rootNodeMock := &mocks.TreeNode{}

	bstContainerMock := &mocks.Container{}
	bstContainerMock.On("GetRoot").Return(rootNodeMock)
	bstContainerMock.On("Find", rootNodeMock, 77).Return("77-value")
	bstContainerMock.On("Delete", rootNodeMock, 77).Return(nil)

	ginMock := &mocks.GinContext{}
	ginMock.On("Query", "val").Return("77")
	ginMock.On("AbortWithStatus", http.StatusOK)

	apiContainer := &Container{
		BstContainer: bstContainerMock,
	}

	apiContainer.DeleteByKey(ginMock)

	bstContainerMock.AssertExpectations(t)
	ginMock.AssertExpectations(t)
}

func TestContainer_DeleteIncorrectValue(t *testing.T) {

	bstContainerMock := &mocks.Container{}

	ginMock := &mocks.GinContext{}
	ginMock.On("Query", "val").Return("incorrect value")
	ginMock.On("AbortWithStatusJSON", http.StatusInternalServerError, mock.Anything)

	apiContainer := &Container{
		BstContainer: bstContainerMock,
	}

	apiContainer.DeleteByKey(ginMock)

	bstContainerMock.AssertExpectations(t)
	ginMock.AssertExpectations(t)
}

func TestContainer_DeleteValueNotFound(t *testing.T) {

	rootNodeMock := &mocks.TreeNode{}

	bstContainerMock := &mocks.Container{}
	bstContainerMock.On("GetRoot").Return(rootNodeMock)
	bstContainerMock.On("Find", rootNodeMock, 77).Return(nil)

	ginMock := &mocks.GinContext{}
	ginMock.On("Query", "val").Return("77")
	ginMock.On("AbortWithStatusJSON", http.StatusNotFound, gin.H{
		"message": "value for key 77 not found",
	})

	apiContainer := &Container{
		BstContainer: bstContainerMock,
	}

	apiContainer.DeleteByKey(ginMock)

	bstContainerMock.AssertExpectations(t)
	ginMock.AssertExpectations(t)
}

func TestContainer_InsertOk(t *testing.T) {

	rootNodeMock := &mocks.TreeNode{}

	bstContainerMock := &mocks.Container{}
	bstContainerMock.On("GetRoot").Return(rootNodeMock)
	bstContainerMock.On("Insert", rootNodeMock, 0, nil).Return(nil)

	ginMock := &mocks.GinContext{}
	ginMock.On("BindJSON", &InsertParams{}).Return(nil)
	ginMock.On("AbortWithStatus", http.StatusOK)

	apiContainer := &Container{
		BstContainer: bstContainerMock,
	}

	apiContainer.Insert(ginMock)

	bstContainerMock.AssertExpectations(t)
	ginMock.AssertExpectations(t)
}

func TestContainer_InsertBindJsonError(t *testing.T) {

	bstContainerMock := &mocks.Container{}

	ginMock := &mocks.GinContext{}
	ginMock.On("BindJSON", &InsertParams{}).Return(errors.New("bind json failed"))
	ginMock.On("AbortWithStatusJSON", http.StatusBadRequest, mock.Anything)

	apiContainer := &Container{
		BstContainer: bstContainerMock,
	}

	apiContainer.Insert(ginMock)

	bstContainerMock.AssertExpectations(t)
	ginMock.AssertExpectations(t)
}

func TestContainer_InsertFailed(t *testing.T) {

	rootNodeMock := &mocks.TreeNode{}

	bstContainerMock := &mocks.Container{}
	bstContainerMock.On("GetRoot").Return(rootNodeMock)
	bstContainerMock.On("Insert", rootNodeMock, 0, nil).Return(errors.New("insert failed"))

	ginMock := &mocks.GinContext{}
	ginMock.On("BindJSON", &InsertParams{}).Return(nil)
	ginMock.On("AbortWithStatusJSON", http.StatusInternalServerError, mock.Anything)

	apiContainer := &Container{
		BstContainer: bstContainerMock,
	}

	apiContainer.Insert(ginMock)

	bstContainerMock.AssertExpectations(t)
	ginMock.AssertExpectations(t)
}
