package sprite

type spriteFrameShared struct {
	collisionMasks [maxCollisionMasks]CollisionMask
}

func (spr *spriteFrameShared) init(frameData spriteAssetFrame) {
	spr.collisionMasks = frameData.CollisionMasks
}
