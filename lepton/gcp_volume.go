package lepton

// CreateVolume is a stub to satisfy VolumeService interface
func (g *GCloud) CreateVolume(config *Config, name, label, data, size, provider string) error {
	return nil
}

// GetAllVolumes is a stub to satisfy VolumeService interface
func (g *GCloud) GetAllVolumes(config *Config) error {
	return nil
}

// UpdateVolume is a stub to satisfy VolumeService interface
func (g *GCloud) UpdateVolume(config *Config, name, label string) error {
	return nil
}

// DeleteVolume is a stub to satisfy VolumeService interface
func (g *GCloud) DeleteVolume(config *Config, name, label string) error {
	return nil
}

// AttachVolume is a stub to satisfy VolumeService interface
func (g *GCloud) AttachVolume(config *Config, image, name, label, mount string) error {
	return nil
}

// DetachVolume is a stub to satisfy VolumeService interface
func (g *GCloud) DetachVolume(config *Config, image, label string) error {
	return nil
}
